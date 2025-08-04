package domain

import "gorm.io/gorm"

type LearningSupply struct {
	gorm.Model
	ProductID uint `gorm:"uniqueIndex"` // One-to-one
	Brand     string
	Material  string
}
