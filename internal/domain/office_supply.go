package domain

import "gorm.io/gorm"

type OfficeSupply struct {
	gorm.Model
	ProductID uint `gorm:"uniqueIndex"` // One-to-one
	Color     string
	Size      string
}
