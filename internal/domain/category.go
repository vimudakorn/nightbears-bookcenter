package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string     `gorm:"unique;not null"`
	ParentID *uint      // Use pointer for nullable foreign key
	Parent   *Category  `gorm:"foreignKey:ParentID" json:"-"` // Hide parent from JSON to prevent infinite loop
	Children []Category `gorm:"foreignKey:ParentID"`
}

type CategoryRepository interface {
	Create(category *Category) error
	FindByID(id uint) (*Category, error)
	FindAll() ([]Category, error)
	FindRootCategories() ([]Category, error)
	FindChildren(parentID uint) ([]Category, error)
	Update(category *Category) error
	Delete(id uint) error
	IsNameExists(name string) (bool, error)
	IsIDExists(id *uint) (bool, error)
	FindAllDescendants(id uint) ([]Category, error)
	FindChildrenRecursive(parentID uint, collected *[]Category) error
}
