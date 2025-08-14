package domain

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID         uint `gorm:"not null"`
	ProductID       *uint
	GroupID         *uint
	Quantity        int     `gorm:"default:1;not null"`
	PriceAtPurchase float64 `gorm:"type:numeric(10,2);not null"`
}

type OrderItemRepository interface {
	Create(item *OrderItem) error
	GetByID(id uint) (*OrderItem, error)
	GetByOrderID(orderID uint) ([]OrderItem, error)
	Update(item *OrderItem) error
	Delete(id uint) error
}
