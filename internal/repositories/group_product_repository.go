package repositories

import (
	"github.com/vimudakorn/internal/domain"
	"gorm.io/gorm"
)

type GroupProductGormRepo struct {
	db *gorm.DB
}

func NewGroupProductGormRepo(db *gorm.DB) domain.GroupProductRepository {
	return &GroupProductGormRepo{db: db}
}

// Create implements domain.GroupProductRepository.
func (g *GroupProductGormRepo) Create(groupProduct *domain.GroupProduct) error {
	return g.db.Create(groupProduct).Error
}

// GetProductByID implements domain.GroupProductRepository.
func (g *GroupProductGormRepo) GetProductByID(groupID uint) ([]domain.GroupProduct, error) {
	panic("unimplemented")
}

func (r *GroupProductGormRepo) CreateWithProduct(tx *gorm.DB, pg *domain.GroupProduct) error {
	return tx.Create(pg).Error
}
