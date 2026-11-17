package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/pgvector/pgvector-go"
	openai "github.com/sashabaranov/go-openai"

	"poc-rag-mcp/backend/internal/controller"
	"poc-rag-mcp/backend/internal/repository"
	"poc-rag-mcp/backend/internal/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system env")
	}

	// 1. Inicializar DB
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	dbPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// 2. Inicializar OpenAI
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("OPENAI_API_KEY is required")
	}
	aiClient := openai.NewClient(apiKey)

	// 3. Procesar embeddings faltantes (RAG)
	processMissingEmbeddings(dbPool, aiClient)

	// 4. Inyección de Dependencias
	orderRepo := repository.NewPostgresOrderRepository(dbPool)
	chatSvc := service.NewChatService(aiClient, orderRepo)
	chatCtrl := controller.NewChatController(chatSvc)

	// 5. Configurar Router Chi
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(CORSMiddleware)

	r.Post("/api/chat", chatCtrl.HandleChat)

	log.Println("🚀 Backend escuchando en http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

// processMissingEmbeddings busca documentos en tienda_docs sin embedding y los genera
func processMissingEmbeddings(db *pgxpool.Pool, ai *openai.Client) {
	rows, err := db.Query(context.Background(), "SELECT id, content FROM tienda_docs WHERE embedding IS NULL")
	if err != nil {
		log.Printf("Error buscando documentos sin embedding: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var content string
		if err := rows.Scan(&id, &content); err != nil {
			continue
		}

		req := openai.EmbeddingRequest{
			Input: []string{content},
			Model: openai.SmallEmbedding3,
		}
		resp, err := ai.CreateEmbeddings(context.Background(), req)
		if err != nil {
			log.Printf("Error generando embedding para doc %d: %v", id, err)
			continue
		}

		emb := pgvector.NewVector(resp.Data[0].Embedding)
		_, err = db.Exec(context.Background(), "UPDATE tienda_docs SET embedding = $1 WHERE id = $2", emb, id)
		if err != nil {
			log.Printf("Error actualizando embedding para doc %d: %v", id, err)
		} else {
			log.Printf("✅ Embedding generado para documento %d", id)
		}
	}
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
