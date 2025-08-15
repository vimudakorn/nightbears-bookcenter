package usecases

import (
	"github.com/vimudakorn/internal/domain"
	orderitemresponse "github.com/vimudakorn/internal/responses/order_item_response"
)

type OrderItemUsecase struct {
	orderItemRepo domain.OrderItemRepository
	orderRepo     domain.OrderRepository
	groupRepo     domain.GroupRepository
	productRepo   domain.ProductRepository
}

func NewOrderItemUsecase(oiRepo domain.OrderItemRepository, oRepo domain.OrderRepository, gRepo domain.GroupRepository, pRepo domain.ProductRepository) *OrderItemUsecase {
	return &OrderItemUsecase{orderItemRepo: oiRepo, orderRepo: oRepo, groupRepo: gRepo, productRepo: pRepo}
}

func (ci *OrderItemUsecase) GetItemsByOrderID(oderID uint) ([]orderitemresponse.OrderItemDetailResponse, error) {
	return ci.orderItemRepo.GetItemsByOrderID(oderID)
}

func (ci *OrderItemUsecase) AddOrUpdateMultOrderItems(orderID uint, items []domain.OrderItem) error {
	return ci.orderItemRepo.AddOrUpdateMultiOrderItems(orderID, items)
}

func (ci *OrderItemUsecase) AddOrUpdateOrderItem(cartID uint, item *domain.OrderItem) error {
	return ci.orderItemRepo.AddOrUpdateOrderItem(cartID, item)
}

func (ci *OrderItemUsecase) GetOrderByUserID(userID uint) ([]domain.Order, error) {
	return ci.orderRepo.GetByUserID(userID)
}

func (ci *OrderItemUsecase) IsGroupIDExist(groupID uint) (bool, error) {
	return ci.groupRepo.IsGroupIDExists(groupID)
}

func (ci *OrderItemUsecase) IsProductIDExists(productID uint) (bool, error) {
	return ci.productRepo.IsProductIDExists(productID)
}

func (ci *OrderItemUsecase) Update(orderItem *domain.OrderItem) error {
	return ci.orderItemRepo.Update(orderItem)
}

func (ci *OrderItemUsecase) UpdateItemInOrderID(orderID uint, orderItems []domain.OrderItem) error {
	return ci.orderItemRepo.UpdateItemsInOrderID(orderID, orderItems)
}

func (ci *OrderItemUsecase) Delete(orderID uint, orderItemID uint) error {
	return ci.orderItemRepo.Delete(orderID, orderItemID)
}
