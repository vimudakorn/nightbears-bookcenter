package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID     uint        `gorm:"not null"`
	User       User        `gorm:"foreignKey:UserID"`
	TotalPrice float64     `gorm:"type:decimal(10,2);not null"`
	Status     string      `gorm:"not null;type:enum('pending','printed','cancelled');default:'pending'"`
	Items      []OrderItem `gorm:"foreignKey:OrderID"`
}
