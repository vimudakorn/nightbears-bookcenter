package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ProductCode   int    `gorm:"unique;not null"`
	ProductType   string `gorm:"not null"`
	Name          string `gorm:"not null"`
	Description   string
	Price         float64 `gorm:"type:numeric(10,2);not null"`
	Discount      float64 `gorm:"type:numeric(10,2);default:0"`
	Stock         int     `gorm:"default:0"`
	ImageURL      string
	CategoryID    *uint `gorm:"not null"`
	Category      Category
	ProductImages []ProductImage `gorm:"foreignKey:ProductID"`
	Tags          []Tag          `gorm:"many2many:product_tags"`

	// ความสัมพันธ์ one-to-one (ไม่ต้องมี ID field ฝั่งนี้)
	Book           *Book           `gorm:"foreignKey:ProductID"`
	LearningSupply *LearningSupply `gorm:"foreignKey:ProductID"`
	OfficeSupply   *OfficeSupply   `gorm:"foreignKey:ProductID"`
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
	FindBookID(productID uint) (uint, error)
	FindLearningID(productID uint) (uint, error)
	FindOfficeID(productID uint) (uint, error)
	IsProductIDExists(productID uint) (bool, error)
}
