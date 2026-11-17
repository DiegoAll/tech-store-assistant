package controller

import (
	"encoding/json"
	"net/http"

	"poc-rag-mcp/backend/internal/service"
)

type ChatController struct {
	chatService service.ChatService
}

func NewChatController(s service.ChatService) *ChatController {
	return &ChatController{chatService: s}
}

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Reply string `json:"reply"`
}

func (c *ChatController) HandleChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ChatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	reply, err := c.chatService.ProcessMessage(r.Context(), req.Message)
	if err != nil {
		http.Error(w, "Error generando respuesta", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ChatResponse{Reply: reply})
}
