package domain

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID         uint `gorm:"not null"`
	ProductID       *uint
	GroupID         *uint
	Quantity        int     `gorm:"default:1;not null"`
	PriceAtPurchase float64 `gorm:"type:decimal(10,2);not null"`
}
