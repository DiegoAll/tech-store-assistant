package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func CheckStockTool() mcp.Tool {
	return mcp.NewTool("check_stock",
		mcp.WithDescription("Consulta la disponibilidad real de un producto en bodega."),
		mcp.WithString("product_name",
			mcp.Required(),
			mcp.Description("El nombre exacto del producto (ej. 'Cámara SecureCam Pro')"),
		),
	)
}

func HandleCheckStock(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.Params.Arguments.(map[string]any)
	productName, _ := args["product_name"].(string)

	var stock int
	query := `SELECT stock FROM products WHERE name ILIKE $1`
	err := dbPool.QueryRow(ctx, query, "%"+productName+"%").Scan(&stock)

	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("El producto '%s' no está en el catálogo.", productName)), nil
	}

	if stock > 0 {
		return mcp.NewToolResultText(fmt.Sprintf("¡Sí! Tenemos %d unidades de '%s' disponibles.", stock, productName)), nil
	}
	return mcp.NewToolResultText(fmt.Sprintf("Lo siento, el producto '%s' está agotado.", productName)), nil
}
