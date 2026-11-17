package repository

import (
	"context"
	"database/sql"
	"errors"

	"poc-rag-mcp/backend/internal/domain"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresOrderRepo struct {
	db *pgxpool.Pool
}

func NewPostgresOrderRepository(db *pgxpool.Pool) domain.OrderRepository {
	return &postgresOrderRepo{db: db}
}

func (r *postgresOrderRepo) GetOrderByID(ctx context.Context, id string) (*domain.Order, error) {
	// Simulamos una consulta a una tabla de pedidos
	// En una app real: SELECT id, status, items FROM orders WHERE id = $1
	query := `SELECT id, status, items FROM orders WHERE id = $1`

	var order domain.Order
	err := r.db.QueryRow(ctx, query, id).Scan(&order.ID, &order.Status, &order.Items)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) || err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &order, nil
}
