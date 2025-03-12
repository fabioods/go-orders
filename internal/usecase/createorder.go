package usecase

import (
	"context"
	"time"

	"github.com/fabioods/go-orders/internal/entity"
	"github.com/fabioods/go-orders/pkg/errorformatted"
	"github.com/fabioods/go-orders/pkg/trace"
)

type (
	CreateOrderUseCase struct {
		UserRepository  UserRepository
		OrderRepository OrderRepository
	}

	CreateOrderInput struct {
		UserID string                 `json:"user_id"`
		Items  []CreateOrderItemInput `json:"items"`
	}

	CreateOrderItemInput struct {
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Quantity int64   `json:"quantity"`
	}

	CreateOrderItemOutput struct {
		ID       string  `json:"id"`
		Name     string  `json:"name"`
		Price    float64 `json:"price"`
		Quantity int64   `json:"quantity"`
	}

	CreateOrderOutput struct {
		ID        string                  `json:"id"`
		UserID    string                  `json:"user_id"`
		Status    string                  `json:"status"`
		Total     float64                 `json:"total"`
		Items     []CreateOrderItemOutput `json:"items"`
		CreatedAt time.Time               `json:"created_at"`
		UpdatedAt time.Time               `json:"updated_at"`
	}
)

//go:generate mockery --name=OrderRepository --output=mocks --case=underscore
type OrderRepository interface {
	Save(ctx context.Context, order *entity.Order) error
}

func NewCreateOrderUseCase(userRepository UserRepository, orderRepository OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		UserRepository:  userRepository,
		OrderRepository: orderRepository,
	}
}

func (c *CreateOrderUseCase) Execute(ctx context.Context, input CreateOrderInput) (*CreateOrderOutput, error) {
	user, err := c.UserRepository.FindByID(ctx, input.UserID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errorformatted.BadRequestError(trace.GetTrace(), "user_not_found", "User not found")
	}

	order := entity.NewOrder(user.ID)
	for _, item := range input.Items {
		orderItem := entity.NewOrderItem(item.Name, item.Price, item.Quantity, order.ID)
		order.AddItem(orderItem)
	}
	order.CalculateTotal()

	if err := c.OrderRepository.Save(ctx, order); err != nil {
		return nil, err
	}

	var items []CreateOrderItemOutput = make([]CreateOrderItemOutput, 0)
	for _, item := range order.Items {
		itemOutput := CreateOrderItemOutput{}
		itemOutput.ID = item.ID
		itemOutput.Name = item.Name
		itemOutput.Price = item.Price
		itemOutput.Quantity = item.Quantity
		items = append(items, itemOutput)
	}

	orderOutput := &CreateOrderOutput{}
	orderOutput.ID = order.ID
	orderOutput.UserID = order.UserID
	orderOutput.Status = string(order.Status)
	orderOutput.Total = order.Total
	orderOutput.Items = items
	orderOutput.CreatedAt = order.CreatedAt
	orderOutput.UpdatedAt = order.UpdatedAt

	return orderOutput, nil
}
