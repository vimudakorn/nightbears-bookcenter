package repositories

import (
	"github.com/vimudakorn/internal/domain"
	"gorm.io/gorm"
)

type OrderItemGormRepo struct {
	db *gorm.DB
}

func (r *OrderItemGormRepo) Create(item *domain.OrderItem) error {
	return r.db.Create(item).Error
}

func (r *OrderItemGormRepo) GetByID(id uint) (*domain.OrderItem, error) {
	var it domain.OrderItem
	if err := r.db.First(&it, id).Error; err != nil {
		return nil, err
	}
	return &it, nil
}

func (r *OrderItemGormRepo) GetByOrderID(orderID uint) ([]domain.OrderItem, error) {
	var items []domain.OrderItem
	if err := r.db.Where("order_id = ?", orderID).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *OrderItemGormRepo) Update(item *domain.OrderItem) error {
	return r.db.Save(item).Error
}

func (r *OrderItemGormRepo) Delete(id uint) error {
	return r.db.Delete(&domain.OrderItem{}, id).Error
}

func NewOrderItemGormRepo(db *gorm.DB) domain.OrderItemRepository {
	return &OrderItemGormRepo{db: db}
}
