package domain

import "gorm.io/gorm"

// type Product struct {
// 	gorm.Model
// 	ProductCode      int    `gorm:"unique;not null"`
// 	ProductType      string `gorm:"not null"`
// 	Name             string `gorm:"not null"`
// 	Description      string
// 	Price            float64 `gorm:"type:decimal(10,2);not null"`
// 	Stock            int     `gorm:"default:0"`
// 	ImageURL         string
// 	BookID           *uint
// 	Book             *Book
// 	LearningSupplyID *uint
// 	LearningSupply   *LearningSupply
// 	OfficeSupplyID   *uint
// 	OfficeSupply     *OfficeSupply
// 	Categories       []Category     `gorm:"many2many:book_categories;"`
// 	BookImages       []ProductImage `gorm:"foreignKey:ProductID"` // ✅ บอกว่า BookImage ใช้ ProductID เป็น FK()
// }

type Product struct {
	gorm.Model
	ProductCode      int    `gorm:"unique;not null"`
	ProductType      string `gorm:"not null"`
	Name             string `gorm:"not null"`
	Description      string
	Price            float64 `gorm:"type:decimal(10,2);not null"`
	Stock            int     `gorm:"default:0"`
	ImageURL         string
	BookID           *uint
	Book             *Book
	LearningSupplyID *uint
	LearningSupply   *LearningSupply
	OfficeSupplyID   *uint
	OfficeSupply     *OfficeSupply
	CategoryID       *uint          `gorm:"not null"` // FK to Category
	Category         Category       // The associated Category object
	ProductImages    []ProductImage `gorm:"foreignKey:ProductID"`
}

type ProductRepository interface {
	CreateProduct(product *Product) error
	CreateBook(book *Book) error
	CreateLearning(learning *LearningSupply) error
	CreateOffice(office *OfficeSupply) error
	FindAll() ([]Product, error)
	GetPagination(page int, limit int, search string, sortBy string, orderBy string) ([]Product, int64, error)
	FindByID(id uint) (*Product, error)
	Update(product *Product) error
	Delete(id uint) error
	DeleteBook(bookID uint) error
	DeleteLearning(learningID uint) error
	DeleteOffice(officeID uint) error
}
