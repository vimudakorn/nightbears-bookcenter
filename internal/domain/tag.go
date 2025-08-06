package domain

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name     string    `gorm:"not null"`
	Products []Product `gorm:"many2many:product_tags"` // reverse relation
}

type TagRepository interface {
	GetTagsByIDs(ids []uint) ([]Tag, error)
	FindAll() ([]Tag, error)
	Create(tag *Tag) error
	Update(tag *Tag) error
	Delete(id uint) error
	GetPagination(page int, limit int, search string, sortBy string, orderBy string) ([]Tag, int64, error)
	RenameTag(tagID uint, newName string) error
	IsTagHasUsed(id uint) (bool, error)
	IsNameExists(name string) (bool, error)
}
