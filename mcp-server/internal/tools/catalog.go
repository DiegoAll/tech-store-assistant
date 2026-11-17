package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

func SearchProductsTool() mcp.Tool {
	return mcp.NewTool("search_products",
		mcp.WithDescription("Busca productos en el catálogo por descripción o características. Devuelve nombre, stock y precio."),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("Palabra clave del producto (ej. 'cámara', 'sensor', 'termostato')"),
		),
		mcp.WithNumber("max_price",
			mcp.Description("Precio máximo en USD para filtrar (opcional)."),
		),
	)
}

func HandleSearchProducts(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.Params.Arguments.(map[string]any)
	query, _ := args["query"].(string)

	// Limpiar la query (quitar signos de interrogación, pasar a minúsculas)
	query = strings.ToLower(strings.Trim(query, "?¿!¡ "))

	// Manejo seguro del precio: puede venir como float64 (número) o string (texto)
	var maxPrice float64
	if val, ok := args["max_price"].(float64); ok {
		maxPrice = val
	} else if valStr, ok := args["max_price"].(string); ok {
		// Si el LLM envía "100", lo convertimos a número de forma segura
		fmt.Sscanf(valStr, "%f", &maxPrice)
	}

	// Buscar por cada palabra clave de la query para hacerlo más flexible
	words := strings.Fields(query)
	if len(words) == 0 {
		return mcp.NewToolResultText("Por favor, ingresa una palabra clave válida."), nil
	}

	// Construimos una query dinámica: busca productos que contengan CUALQUIERA de las palabras
	sqlQuery := `SELECT name, stock, price FROM products WHERE (`
	sqlArgs := []any{}
	for i, word := range words {
		if i > 0 {
			sqlQuery += " OR "
		}
		sqlQuery += fmt.Sprintf("name ILIKE $%d", i+1)
		sqlArgs = append(sqlArgs, "%"+word+"%")
	}
	sqlQuery += `)`

	// Aplicar filtro de precio si existe
	if maxPrice > 0 {
		sqlQuery += fmt.Sprintf(" AND price <= $%d", len(sqlArgs)+1)
		sqlArgs = append(sqlArgs, maxPrice)
	}

	sqlQuery += ` ORDER BY price ASC LIMIT 3`

	// Ejecutar consulta
	rows, err := dbPool.Query(ctx, sqlQuery, sqlArgs...)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error en la búsqueda: %v", err)), nil
	}
	defer rows.Close()

	var result string
	count := 0
	for rows.Next() {
		var name string
		var stock int
		var price float64
		if err := rows.Scan(&name, &stock, &price); err == nil {
			result += fmt.Sprintf("- %s (Precio: $%.2f, Stock: %d)\n", name, price, stock)
			count++
		}
	}

	if count == 0 {
		return mcp.NewToolResultText(fmt.Sprintf("No encontré productos que coincidan con '%s'.", query)), nil
	}

	return mcp.NewToolResultText("Encontré estos productos:\n" + result), nil
}
