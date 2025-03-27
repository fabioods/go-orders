package usecase

import "context"

type ProcessOrderInput struct {
	OrderID string `json:"id"`
}

type ProcessOrderUseCase struct {
	OrderRepository OrderRepository
}

func NewProcessOrderUseCase(
	orderRepository OrderRepository,
) *ProcessOrderUseCase {
	return &ProcessOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (u *ProcessOrderUseCase) Execute(ctx context.Context, input ProcessOrderInput) error {
	order, err := u.OrderRepository.FindByID(ctx, input.OrderID)
	if err != nil {
		return err
	}

	if order == nil {
		return nil
	}

	order.Status = "PROCESSED"
	err = u.OrderRepository.UpdateProcessOrder(ctx, order)
	if err != nil {
		return err
	}

	return nil
}
