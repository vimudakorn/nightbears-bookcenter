package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID     uint        `gorm:"not null"`
	User       User        `gorm:"foreignKey:UserID"`
	TotalPrice float64     `gorm:"type:numeric(10,2);not null"`
	Status     string      `gorm:"not null"`
	Items      []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;"`
}

type OrderRepository interface {
	Create(order *Order) error
	GetByID(id uint) (*Order, error)
	GetByUserID(userID uint) ([]Order, error)
	GetAll(page, limit int, search, sortBy, orderBy string) ([]Order, int64, error)
	Update(order *Order) error
	Delete(id uint) error
	UpdateOrderFields(orderID uint, fields map[string]interface{}) error
	UpdateItemsInOrderID(orderID uint, items []OrderItem) error
}
