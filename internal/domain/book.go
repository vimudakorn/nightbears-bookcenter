package domain

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string
	Author      string
	ISBN        string
	Description string
	Price       float64
	Discount    float64    // Consider adding a validation that Discount <= Price
	ImageURL    string     `gorm:"column:image_url"`           // Use explicit column name for snake_case
	Categories  []Category `gorm:"many2many:book_categories;"` // Removed duplicate 'many2many' tag
}
