package domain

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID uint       `gorm:"unique;not null"`
	User   *User      `gorm:"foreignKey:UserID" json:"-"`
	Items  []CartItem `gorm:"foreignKey:CartID;constraint:OnDelete:CASCADE;"`
}

type CartRepository interface {
	Create(cart *Cart) error
	GetByUserID(userID uint) (*Cart, error)
	Update(cart *Cart) error
	Delete(cartID uint) error
}
