package domain

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	CartID    uint `gorm:"not null"`
	ProductID *uint
	GroupID   *uint
	Quantity  int `gorm:"default:1;not null"`
}
