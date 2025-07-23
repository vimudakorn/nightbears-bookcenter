package domain

import "gorm.io/gorm"

type BookCategory struct {
	gorm.Model
	BookID     uint
	CategoryID uint
}
