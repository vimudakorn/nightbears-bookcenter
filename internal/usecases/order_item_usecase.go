package usecases

import "github.com/vimudakorn/internal/domain"

type OrderItemUsecase struct {
	orderItemRepo domain.OrderItemRepository
}

func NewOrderItemUsecase(oiRepo domain.OrderItemRepository) *OrderItemUsecase {
	return &OrderItemUsecase{orderItemRepo: oiRepo}
}
