package handlers

import "github.com/vimudakorn/internal/usecases"

type OrderItemHandler struct {
	usecases *usecases.OrderItemUsecase
}

func NewOrderItemHandler(uc *usecases.OrderItemUsecase) *OrderItemHandler {
	return &OrderItemHandler{usecases: uc}
}
