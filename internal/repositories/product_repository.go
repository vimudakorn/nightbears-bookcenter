package repositories

import (
	"fmt"

	"github.com/vimudakorn/internal/domain"
	productrequest "github.com/vimudakorn/internal/request/product_request"
	"gorm.io/gorm"
)

type ProductGormRepo struct {
	db *gorm.DB
}

// IsProductIDExists implements domain.ProductRepository.
func (b *ProductGormRepo) IsProductIDExists(productID uint) (bool, error) {
	var count int64
	err := b.db.Model(&domain.Product{}).Where("id = ?", productID).Count(&count).Error
	return count > 0, err

}

// DeleteBook implements domain.ProductRepository.
func (b *ProductGormRepo) DeleteBook(bookID uint) error {
	return b.db.Delete(&domain.Book{}, bookID).Error
}

// DeleteLearning implements domain.ProductRepository.
func (b *ProductGormRepo) DeleteLearning(learningID uint) error {
	return b.db.Delete(&domain.LearningSupply{}, learningID).Error
}

// DeleteOffice implements domain.ProductRepository.
func (b *ProductGormRepo) DeleteOffice(officeID uint) error {
	return b.db.Delete(&domain.OfficeSupply{}, officeID).Error
}

// CreateBook implements domain.ProductRepository.
func (b *ProductGormRepo) CreateBook(book *domain.Book) error {
	return b.db.Create(book).Error
}

// CreateLearning implements domain.ProductRepository.
func (b *ProductGormRepo) CreateLearning(learning *domain.LearningSupply) error {
	return b.db.Create(learning).Error
}

// CreateOffice implements domain.ProductRepository.
func (b *ProductGormRepo) CreateOffice(office *domain.OfficeSupply) error {
	return b.db.Create(office).Error
}

func NewProductGormDB(db *gorm.DB) domain.ProductRepository {
	return &ProductGormRepo{db: db}
}

func (b *ProductGormRepo) CreateProduct(Product *domain.Product) error {
	return b.db.Create(Product).Error
}

// Delete implements domain.ProductRepository.
func (b *ProductGormRepo) Delete(id uint) error {
	return b.db.Delete(&domain.Product{}, id).Error
}

// FindAll implements domain.ProductRepository.
func (b *ProductGormRepo) FindAll() ([]domain.Product, error) {
	Products := []domain.Product{}
	err := b.db.
		Preload("Category").
		Preload("Tags").
		Preload("Book").
		Preload("LearningSupply").
		Preload("OfficeSupply").Find(&Products).Error
	return Products, err
}

// FindByID implements domain.ProductRepository.
func (b *ProductGormRepo) FindByID(id uint) (*domain.Product, error) {
	Product := domain.Product{}
	err := b.db.
		Preload("Category").
		Preload("Tags").
		Preload("Book").
		Preload("LearningSupply").
		Preload("OfficeSupply").First(&Product, id).Error
	if err != nil {
		return nil, err
	}
	return &Product, nil
}

func (b *ProductGormRepo) GetFilteredProducts(page, limit int, sortBy, orderBy string, filter productrequest.ProductFilter) ([]domain.Product, int64, error) {
	var products []domain.Product
	var count int64

	offset := (page - 1) * limit
	order := fmt.Sprintf("%s %s", sortBy, orderBy)

	query := b.db.Model(&domain.Product{})

	// Dynamic filters
	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}

	if filter.ProductType != "" {
		query = query.Where("product_type = ?", filter.ProductType)
	}

	// Filter fields เฉพาะ book
	if filter.ProductType == "book" {
		if filter.Subject != "" {
			query = query.Joins("LEFT JOIN books b ON b.product_id = products.id").
				Where("b.subject ILIKE ?", "%"+filter.Subject+"%")
		}
		if filter.Grade != "" {
			query = query.Joins("LEFT JOIN books b ON b.product_id = products.id").
				Where("b.grade = ?", filter.Grade)
		}
		if filter.Publisher != "" {
			query = query.Joins("LEFT JOIN books b ON b.product_id = products.id").
				Where("b.publisher ILIKE ?", "%"+filter.Publisher+"%")
		}
	}

	query.Count(&count)

	err := query.
		Preload("Category").
		Preload("Tags").
		Preload("Book").
		Preload("LearningSupply").
		Preload("OfficeSupply").
		Order(order).
		Limit(limit).
		Offset(offset).
		Find(&products).Error

	return products, count, err
}

// GetPagination implements domain.ProductRepository.
func (b *ProductGormRepo) GetPagination(page int, limit int, search string, sortBy string, orderBy string) ([]domain.Product, int64, error) {
	var products []domain.Product
	var count int64

	allowedSortBy := map[string]bool{
		"name": true,
	}
	allowedOrderBy := map[string]bool{
		"asc":  true,
		"desc": true,
	}

	if !allowedSortBy[sortBy] {
		sortBy = "id"
	}
	if !allowedOrderBy[orderBy] {
		orderBy = "asc"
	}

	offset := (page - 1) * limit
	order := fmt.Sprintf("%s %s", sortBy, orderBy)

	query := b.db.Model(&domain.Product{})

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	query.Count(&count)

	err := query.
		Preload("Category").
		Preload("Tags").
		Preload("Book").
		Preload("LearningSupply").
		Preload("OfficeSupply").Order(order).Limit(limit).Offset(offset).Find(&products).Error
	return products, count, err
}

// Update implements domain.ProductRepository.
func (b *ProductGormRepo) Update(product *domain.Product) error {
	return b.db.Save(product).Error
}

func (b *ProductGormRepo) FindBookID(productID uint) (uint, error) {
	var book domain.Book
	err := b.db.Where("product_id = ?", productID).First(&book).Error
	return book.ID, err
}

func (b *ProductGormRepo) FindLearningID(productID uint) (uint, error) {
	var learning domain.LearningSupply
	err := b.db.Where("product_id = ?", productID).First(&learning).Error
	return learning.ID, err
}

func (b *ProductGormRepo) FindOfficeID(productID uint) (uint, error) {
	var office domain.OfficeSupply
	err := b.db.Where("product_id = ?", productID).First(&office).Error
	return office.ID, err
}
