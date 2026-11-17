package tools

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// dbPool es la conexión global compartida por todas las herramientas
var dbPool *pgxpool.Pool

// InitDB inicializa la conexión a PostgreSQL
func InitDB() error {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	var err error
	dbPool, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return fmt.Errorf("error conectando a la base de datos: %v", err)
	}

	// Probar la conexión
	if err := dbPool.Ping(context.Background()); err != nil {
		return fmt.Errorf("error haciendo ping a la base de datos: %v", err)
	}

	log.Println("✅ MCP conectado a PostgreSQL")
	return nil
}
