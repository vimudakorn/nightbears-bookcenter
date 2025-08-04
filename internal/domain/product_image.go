package domain

import (
	"gorm.io/gorm"
)

type ProductImage struct {
	gorm.Model
	ProductID uint    `json:"book_id" gorm:"not null"`
	Product   Product `json:"-"` // Add Book relation, hide from JSON to prevent infinite loop
	ImageURL  string  `json:"image_url" gorm:"not null"`
	IsMain    bool    `json:"is_main" gorm:"default:false"` // Indicates if it's the main image
	SortOrder int     `json:"sort_order"`                   // For ordering images
}
