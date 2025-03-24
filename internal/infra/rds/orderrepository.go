package rds

import (
	"context"

	"github.com/fabioods/go-orders/internal/entity"
	"github.com/fabioods/go-orders/internal/errorcode"
	"github.com/fabioods/go-orders/pkg/errorformatted"
	"github.com/fabioods/go-orders/pkg/trace"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepositoryRDS(db *gorm.DB) *OrderRepository {
	db.AutoMigrate(&entity.Order{}, &entity.OrderItem{})
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) Save(ctx context.Context, order *entity.Order) error {
	err := r.db.Omit("Items.*.ID").Create(order).Error
	if err != nil {
		errFmt := errorformatted.UnexpectedError(trace.GetTrace(), errorcode.ErrorSaveOrderError, "Failed to save Order %s", err.Error())
		return errFmt
	}
	return nil
}
