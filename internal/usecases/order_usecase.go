package usecases

import (
	"github.com/vimudakorn/internal/domain"
	orderitemresponse "github.com/vimudakorn/internal/responses/order_item_response"
)

type OrderUsecase struct {
	orderRepo     domain.OrderRepository
	orderItemRepo domain.OrderItemRepository
}

func NewOrderUsecase(oRepo domain.OrderRepository, oiRepoo domain.OrderItemRepository) *OrderUsecase {
	return &OrderUsecase{orderRepo: oRepo, orderItemRepo: oiRepoo}
}

func (uc *OrderUsecase) CreateOrder(o *domain.Order) error {
	var total float64
	for _, it := range o.Items {
		total += float64(it.Quantity) * it.PriceAtPurchase
	}
	o.TotalPrice = total
	if o.Status == "" {
		o.Status = "PENDING"
	}
	return uc.orderRepo.Create(o)
}
func (uc *OrderUsecase) GetOryderByUserID(userID uint) ([]domain.Order, error) {
	return uc.orderRepo.GetByUserID(userID)
}

func (uc *OrderUsecase) GetOrderByID(id uint) (*domain.Order, error) {
	return uc.orderRepo.GetByID(id)
}
func (uc *OrderUsecase) GetOrders(page, limit int, search, sortBy, orderBy string) ([]domain.Order, int64, error) {
	return uc.orderRepo.GetAll(page, limit, search, sortBy, orderBy)
}
func (uc *OrderUsecase) UpdateOrder(o *domain.Order) error {
	return uc.orderRepo.Update(o)
}
func (uc *OrderUsecase) DeleteOrder(id uint) error {
	return uc.orderRepo.Delete(id)
}

// Items
func (uc *OrderUsecase) AddOrderItem(orderID uint, it *domain.OrderItem) error {
	return uc.orderItemRepo.AddOrUpdateOrderItem(orderID, it)
}
func (uc *OrderUsecase) GetOrderItems(orderID uint) ([]orderitemresponse.OrderItemDetailResponse, error) {
	return uc.orderItemRepo.GetItemsByOrderID(orderID)
}
func (uc *OrderUsecase) UpdateOrderItem(it *domain.OrderItem) error {
	return uc.orderItemRepo.Update(it)
}
func (uc *OrderUsecase) DeleteOrderItem(orderID uint, orderItemID uint) error {
	return uc.orderItemRepo.Delete(orderID, orderItemID)
}
