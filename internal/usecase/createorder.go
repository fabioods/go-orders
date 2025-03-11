package usecase

import (
	"context"

	"github.com/fabioods/go-orders/internal/entity"
)

type (
	CreateOrderItemDTO struct {
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Quantity int64   `json:"quantity"`
	}

	CreateOrderDTO struct {
		UserID string               `json:"user_id"`
		Status string               `json:"status"`
		Items  []CreateOrderItemDTO `json:"items"`
	}

	CreateOrderUseCase struct {
		UserRepository  UserRepository
		OrderRepository OrderRepository
	}

	OrderRepository interface {
		Save(ctx context.Context, order *entity.Order) error
	}
)

func NewCreateOrderUseCase(userRepository UserRepository, orderRepository OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		UserRepository:  userRepository,
		OrderRepository: orderRepository,
	}
}

func (c *CreateOrderUseCase) Execute(ctx context.Context, input CreateOrderDTO) (*entity.Order, error) {
	user, err := c.UserRepository.FindByID(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	order := entity.NewOrder(user.ID)
	for _, item := range input.Items {
		orderItem := &entity.OrderItem{
			Name:     item.Name,
			Price:    item.Price,
			Quantity: item.Quantity,
		}
		order.AddItem(orderItem)
	}

	if err := c.OrderRepository.Save(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}
