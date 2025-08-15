package repositories

import (
	"fmt"

	"github.com/vimudakorn/internal/domain"
	"gorm.io/gorm"
)

type OrderGormRepo struct {
	db *gorm.DB
}

// GetByUserID implements domain.OrderRepository.
func (r *OrderGormRepo) GetByUserID(userID uint) ([]domain.Order, error) {
	var orders []domain.Order
	if err := r.db.Preload("Items").
		Where("user_id = ?", userID).
		Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderGormRepo) Create(order *domain.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderGormRepo) GetByID(id uint) (*domain.Order, error) {
	var o domain.Order
	if err := r.db.Preload("Items").First(&o, id).Error; err != nil {
		return nil, err
	}
	return &o, nil
}

func (r *OrderGormRepo) GetAll(page, limit int, search, sortBy, orderBy string) ([]domain.Order, int64, error) {
	if sortBy == "" {
		sortBy = "id"
	}
	if orderBy == "" {
		orderBy = "asc"
	}
	offset := (page - 1) * limit

	q := r.db.Model(&domain.Order{})
	if search != "" {
		q = q.Where("CAST(user_id AS TEXT) ILIKE ? OR status ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	var total int64
	q.Count(&total)

	var orders []domain.Order
	if err := q.Preload("Items").Order(fmt.Sprintf("%s %s", sortBy, orderBy)).Limit(limit).Offset(offset).Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	return orders, total, nil
}

func (r *OrderGormRepo) Update(order *domain.Order) error {
	return r.db.Save(order).Error
}

func (r *OrderGormRepo) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("order_id = ?", id).Delete(&domain.OrderItem{}).Error; err != nil {
			return err
		}
		return tx.Delete(&domain.Order{}, id).Error
	})
}

func NewOrderGormRepo(db *gorm.DB) domain.OrderRepository {
	return &OrderGormRepo{db: db}
}
