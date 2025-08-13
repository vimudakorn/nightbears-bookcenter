package repositories

import (
	"github.com/vimudakorn/internal/domain"
	"gorm.io/gorm"
)

type CartGormRepo struct {
	db *gorm.DB
}

// Delete implements domain.CartRepository.
func (c *CartGormRepo) Delete(cartID uint) error {
	return c.db.Delete(&domain.Cart{}, cartID).Error
}

// GetByUserID implements domain.CartRepository.
func (c *CartGormRepo) GetByUserID(userID uint) (*domain.Cart, error) {
	var cart domain.Cart
	err := c.db.Preload("Items").Where("user_id = ?", userID).First(&cart).Error
	return &cart, err
}

// Update implements domain.CartRepository.
func (c *CartGormRepo) Update(cart *domain.Cart) error {
	return c.db.Save(cart).Error
}

func (c *CartGormRepo) Create(cart *domain.Cart) error {
	return c.db.Create(cart).Error
}

func NewCartGormRepo(db *gorm.DB) domain.CartRepository {
	return &CartGormRepo{db: db}
}
