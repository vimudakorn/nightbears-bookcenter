package domain

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID         uint
	Order           Order `json:"-"` // Add Order relation, hide from JSON to prevent infinite loop
	BookID          uint
	Book            Book // Add Book relation
	Quantity        int
	PriceAtPurchase float64
}
