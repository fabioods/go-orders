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

func (r *OrderRepository) FindByID(ctx context.Context, id string) (*entity.Order, error) {
	var order entity.Order
	err := r.db.WithContext(ctx).Preload("Items").Where("id = ?", id).First(&order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		errFmt := errorformatted.UnexpectedError(trace.GetTrace(), errorcode.ErrorFindOrderError, "Failed to find Order %s", err.Error())
		return nil, errFmt
	}
	return &order, nil
}

func (r *OrderRepository) UpdateProcessOrder(ctx context.Context, order *entity.Order) error {
	err := r.db.WithContext(ctx).Model(order).Updates(order).Error
	if err != nil {
		errFmt := errorformatted.UnexpectedError(trace.GetTrace(), errorcode.ErrorUpdateOrderError, "Failed to update Order %s", err.Error())
		return errFmt
	}
	return nil
}
