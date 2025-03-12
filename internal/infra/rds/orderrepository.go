package rds

import (
	"context"

	"github.com/fabioods/go-orders/internal/entity"
)

type OrderRepository struct{}

func NewOrderRepositoryRDS() *OrderRepository {
	return &OrderRepository{}
}

func (r *OrderRepository) Save(ctx context.Context, order *entity.Order) error {
	return nil
}
