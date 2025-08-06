package domain

import "gorm.io/gorm"

type GroupProduct struct {
	gorm.Model
	GroupID   uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	Quantity  int  `gorm:"not null"`
}

type GroupProductRepository interface {
	Create(groupProduct *GroupProduct) error
}
