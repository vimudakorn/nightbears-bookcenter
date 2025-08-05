package domain

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ProductID uint `gorm:"uniqueIndex"` // One-to-one
	Author    string
	ISBN      string
}
