package domain

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name     string    `gorm:"not null"`
	Products []Product `gorm:"many2many:product_tags"` // reverse relation
}

type TagRepository interface {
	GetTagsByIDs(ids []uint) ([]Tag, error)
}
