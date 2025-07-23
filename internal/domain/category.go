package domain

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string
	ParentID *uint      // Use pointer for nullable foreign key
	Parent   *Category  `gorm:"foreignKey:ParentID" json:"-"` // Hide parent from JSON to prevent infinite loop
	Children []Category `gorm:"foreignKey:ParentID"`
}
