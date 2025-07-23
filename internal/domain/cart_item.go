package domain

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID   uint
	Cart     Cart `json:"-"` // Add Cart relation, hide from JSON
	BookID   uint
	Book     Book // Add Book relation
	Quantity uint
}
