package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func CreateTicketTool() mcp.Tool {
	return mcp.NewTool("create_support_ticket",
		mcp.WithDescription("Crea un ticket de soporte técnico para un problema del cliente."),
		mcp.WithString("issue",
			mcp.Required(),
			mcp.Description("Descripción del problema (ej. 'La cámara no conecta al WiFi')"),
		),
	)
}

func HandleCreateTicket(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	args := request.Params.Arguments.(map[string]any)
	issue, _ := args["issue"].(string)

	// En una app real, aquí obtendrías el userID de la sesión
	userID := "user_123"

	var ticketID int
	query := `INSERT INTO support_tickets (user_id, issue) VALUES ($1, $2) RETURNING id`
	err := dbPool.QueryRow(ctx, query, userID, issue).Scan(&ticketID)

	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error al crear el ticket: %v", err)), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("He creado el ticket de soporte #%d para tu problema. Un técnico te contactará pronto.", ticketID)), nil
}
