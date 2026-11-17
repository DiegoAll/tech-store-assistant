package main

import (
	"log"

	"mcp-server/internal/tools"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// 1. Inicializar conexión a la base de datos
	if err := tools.InitDB(); err != nil {
		log.Fatalf("Error inicializando DB: %v", err)
	}

	// 2. Crear el servidor MCP
	s := server.NewMCPServer(
		"TechStore Agent MCP",
		"1.1.0",
		server.WithLogging(),
	)

	// 3. Registrar todas las herramientas
	// Caso 1: Seguimiento de pedidos
	s.AddTool(tools.GetOrderStatusTool(), tools.HandleGetOrderStatus)

	// Caso 2: Stock y reservas
	s.AddTool(tools.CheckStockTool(), tools.HandleCheckStock)

	// Caso 3: Soporte / Troubleshooting
	s.AddTool(tools.CreateTicketTool(), tools.HandleCreateTicket)

	// Caso 5: Búsqueda semántica y comparación
	s.AddTool(tools.SearchProductsTool(), tools.HandleSearchProducts)

	// 4. Configurar el servidor HTTP para SSE
	sseServer := server.NewSSEServer(s,
		server.WithSSEEndpoint("/sse"),
		server.WithMessageEndpoint("/message"),
	)

	log.Println("🚀 Servidor MCP escuchando en http://localhost:5000")
	if err := sseServer.Start(":5000"); err != nil {
		log.Fatalf("Error del servidor MCP: %v", err)
	}
}
