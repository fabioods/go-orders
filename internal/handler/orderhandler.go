package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/fabioods/go-orders/internal/errorcode"
	"github.com/fabioods/go-orders/internal/infra/webserver"
	"github.com/fabioods/go-orders/internal/usecase"
	"github.com/fabioods/go-orders/pkg/errorformatted"
	"github.com/fabioods/go-orders/pkg/response"
	"github.com/fabioods/go-orders/pkg/trace"
)

type (
	OrderHandler struct {
		CreateOrderUseCase CreateOrderUseCase
	}
)

type CreateOrderUseCase interface {
	Execute(ctx context.Context, input usecase.CreateOrderInput) (*usecase.CreateOrderOutput, error)
}

func NewOrderHandler(createOrderUseCase CreateOrderUseCase) *OrderHandler {
	return &OrderHandler{
		CreateOrderUseCase: createOrderUseCase,
	}
}

func (h *OrderHandler) AddOrderHandler(web *webserver.WebServer) {
	web.AddRoute(http.MethodPost, "/orders", h.AddOrder)
}

func (h *OrderHandler) AddOrder(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateOrderInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		ef := errorformatted.BadRequestError(trace.GetTrace(), errorcode.ErrorAvatarFileError, "%s", err.Error())
		response.WriteResponse(w, nil, ef, http.StatusBadRequest)
	}

	orderCreated, err := h.CreateOrderUseCase.Execute(r.Context(), input)

	response.WriteResponse(w, orderCreated, err, http.StatusCreated)
}
