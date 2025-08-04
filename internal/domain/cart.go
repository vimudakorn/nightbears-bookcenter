package domain

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID uint       `gorm:"unique;not null"`
	User   *User      `gorm:"foreignKey:UserID" json:"-"` // Add User relation, hide from JSON
	Items  []CartItem `gorm:"foreignKey:CartID"`
}
