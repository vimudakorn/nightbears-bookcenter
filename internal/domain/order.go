package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID     uint
	User       User `gorm:"foreignKey:UserID"`
	TotalPrice float64
	Status     string
	Items      []OrderItem `gorm:"foreignKey:OrderID"`
}
