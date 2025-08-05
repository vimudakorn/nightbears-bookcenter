package repositories

import (
	"fmt"

	"github.com/vimudakorn/internal/domain"
	"gorm.io/gorm"
)

type ProductGormRepo struct {
	db *gorm.DB
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

// Create implements domain.ProductRepository.
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
	err := b.db.Find(&Products).Error
	return Products, err
}

// FindByID implements domain.ProductRepository.
func (b *ProductGormRepo) FindByID(id uint) (*domain.Product, error) {
	Product := domain.Product{}
	err := b.db.First(&Product, id).Error
	if err != nil {
		return nil, err
	}
	return &Product, nil
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

	err := query.Order(order).Limit(limit).Offset(offset).Find(&products).Error
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
