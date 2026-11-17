package domain

import "context"

type Order struct {
	ID     string
	Status string
	Items  string
}

// OrderRepository es la interfaz abstracta (puerto)
type OrderRepository interface {
	GetOrderByID(ctx context.Context, id string) (*Order, error)
}
