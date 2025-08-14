package usecases

import "github.com/vimudakorn/internal/domain"

type OrderUsecase struct {
	orderRepo     domain.OrderRepository
	orderItemRepo domain.OrderItemRepository
}

func NewOrderUsecase(oRepo domain.OrderRepository, oiRepoo domain.OrderItemRepository) *OrderUsecase {
	return &OrderUsecase{orderRepo: oRepo, orderItemRepo: oiRepoo}
}

func (uc *OrderUsecase) CreateOrder(o *domain.Order) error {
	// auto total
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
func (uc *OrderUsecase) AddOrderItem(it *domain.OrderItem) error {
	return uc.orderItemRepo.Create(it)
}
func (uc *OrderUsecase) GetOrderItems(orderID uint) ([]domain.OrderItem, error) {
	return uc.orderItemRepo.GetByOrderID(orderID)
}
func (uc *OrderUsecase) UpdateOrderItem(it *domain.OrderItem) error {
	return uc.orderItemRepo.Update(it)
}
func (uc *OrderUsecase) DeleteOrderItem(id uint) error {
	return uc.orderItemRepo.Delete(id)
}
