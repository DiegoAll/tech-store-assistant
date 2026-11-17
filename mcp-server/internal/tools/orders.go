package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func GetOrderStatusTool() mcp.Tool {
	return mcp.NewTool("get_order_status",
		mcp.WithDescription("Consulta el estado actual de un pedido en la tienda."),
		mcp.WithString("order_id",
			mcp.Required(),
			mcp.Description("El ID del pedido a consultar (ej. 1234)"),
		),
	)
}

func HandleGetOrderStatus(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.Params.Arguments.(map[string]any)
	orderID, ok := args["order_id"].(string)
	if !ok || orderID == "" {
		return mcp.NewToolResultError("El order_id es requerido"), nil
	}

	if dbPool == nil {
		return mcp.NewToolResultError("Base de datos no inicializada"), nil
	}

	var status string
	query := `SELECT status FROM orders WHERE id = $1`
	err := dbPool.QueryRow(ctx, query, orderID).Scan(&status)

	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Pedido %s no encontrado en el sistema.", orderID)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("El estado del pedido %s es: %s", orderID, status)), nil
}
