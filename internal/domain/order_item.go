package domain

import (
	orderitemresponse "github.com/vimudakorn/internal/responses/order_item_response"
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	OrderID         uint `gorm:"not null"`
	ProductID       *uint
	GroupID         *uint
	Quantity        int     `gorm:"default:1;not null"`
	PriceAtPurchase float64 `gorm:"type:numeric(10,2);not null"`
}

type OrderItemRepository interface {
	Delete(orderID uint, orderItemID uint) error
	Update(item *OrderItem) error
	UpdateItemsInOrderID(orderID uint, items []OrderItem) error
	GetItemsByOrderID(orderID uint) ([]orderitemresponse.OrderItemDetailResponse, error)
	AddOrUpdateOrderItem(orderID uint, item *OrderItem) error
	AddOrUpdateMultiOrderItems(orderID uint, items []OrderItem) error
}
