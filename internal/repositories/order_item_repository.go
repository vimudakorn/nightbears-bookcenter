package repositories

import (
	"errors"
	"fmt"

	"github.com/vimudakorn/internal/domain"
	orderitemresponse "github.com/vimudakorn/internal/responses/order_item_response"
	"gorm.io/gorm"
)

type OrderItemGormRepo struct {
	db *gorm.DB
}

func (r *OrderItemGormRepo) Delete(orderID uint, orderItemID uint) error {
	var order domain.Order
	if err := r.db.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("order with ID %d does not exist", orderID)
		}
		return err
	}

	var orderItem domain.OrderItem
	if err := r.db.Where("id = ? AND order_id = ?", orderItemID, orderID).First(&orderItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("order item with ID %d not found in order %d", orderItemID, orderID)
		}
		return err
	}

	if err := r.db.Delete(&orderItem).Error; err != nil {
		return err
	}

	return nil
}

func (r *OrderItemGormRepo) Update(item *domain.OrderItem) error {
	query := r.db.Model(&domain.OrderItem{}).Where("order_id = ?", item.OrderID)

	if item.ProductID != nil {
		query = query.Where("product_id = ?", *item.ProductID)
	} else {
		query = query.Where("product_id IS NULL")
	}

	if item.GroupID != nil {
		query = query.Where("group_id = ?", *item.GroupID)
	} else {
		query = query.Where("group_id IS NULL")
	}

	return query.Updates(map[string]interface{}{
		"quantity": item.Quantity,
	}).Error
}

func (r *OrderItemGormRepo) UpdateItemsInOrderID(orderID uint, items []domain.OrderItem) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range items {
			query := tx.Model(&domain.OrderItem{}).Where("order_id = ?", orderID)

			if item.ProductID != nil {
				query = query.Where("product_id = ?", *item.ProductID)
			} else {
				query = query.Where("product_id IS NULL")
			}

			if item.GroupID != nil {
				query = query.Where("group_id = ?", *item.GroupID)
			} else {
				query = query.Where("group_id IS NULL")
			}

			var existing domain.OrderItem
			err := query.First(&existing).Error

			if errors.Is(err, gorm.ErrRecordNotFound) {
				continue
			}
			if err != nil {
				return err
			}

			existing.Quantity = item.Quantity
			if err := tx.Save(&existing).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *OrderItemGormRepo) GetItemsByOrderID(orderID uint) ([]orderitemresponse.OrderItemDetailResponse, error) {
	var result []orderitemresponse.OrderItemDetailResponse

	err := r.db.Table("order_items oi").
		Select(`
			oi.order_id,
			u.id AS user_id,
			u.email AS user_email,
			oi.product_id,
			p.name AS product_name,
			p.price AS product_price,
			oi.group_id,
			g.name AS group_name,
			g.sale_price AS group_price,
			oi.quantity
		`).
		Joins("JOIN orders o ON o.id = oi.order_id AND o.deleted_at IS NULL").
		Joins("JOIN users u ON u.id = o.user_id").
		Joins("LEFT JOIN products p ON p.id = oi.product_id AND p.deleted_at IS NULL").
		Joins("LEFT JOIN groups g ON g.id = oi.group_id AND g.deleted_at IS NULL").
		Where("oi.order_id = ? AND oi.deleted_at IS NULL", orderID).
		Scan(&result).Error

	return result, err
}

func (r *OrderItemGormRepo) AddOrUpdateOrderItem(orderID uint, item *domain.OrderItem) error {
	var order domain.Order
	if err := r.db.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("order with ID %d does not exist", orderID)
		}
		return err
	}

	if (item.ProductID != nil && item.GroupID != nil) || (item.ProductID == nil && item.GroupID == nil) {
		return errors.New("either product_id or group_id must be set, but not both")
	}

	var existing domain.OrderItem
	var err error

	if item.ProductID != nil {
		err = r.db.Where("order_id = ? AND product_id = ?", orderID, *item.ProductID).
			First(&existing).Error
	} else {
		err = r.db.Where("order_id = ? AND group_id = ?", orderID, *item.GroupID).
			First(&existing).Error
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if existing.ID != 0 {
		existing.Quantity += item.Quantity
		return r.db.Save(&existing).Error
	}

	item.OrderID = orderID
	return r.db.Create(item).Error
}

func (r *OrderItemGormRepo) AddOrUpdateMultiOrderItems(orderID uint, items []domain.OrderItem) error {
	var order domain.Order
	if err := r.db.First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("order with ID %d does not exist", orderID)
		}
		return err
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range items {
			if (item.ProductID != nil && item.GroupID != nil) || (item.ProductID == nil && item.GroupID == nil) {
				return errors.New("either product_id or group_id must be set, but not both")
			}

			var existing domain.OrderItem
			var err error

			if item.ProductID != nil {
				err = tx.Where("order_id = ? AND product_id = ?", orderID, *item.ProductID).First(&existing).Error
			} else {
				err = tx.Where("order_id = ? AND group_id = ?", orderID, *item.GroupID).First(&existing).Error
			}

			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			if existing.ID != 0 {
				existing.Quantity += item.Quantity
				if err := tx.Save(&existing).Error; err != nil {
					return err
				}
			} else {
				item.OrderID = orderID
				if err := tx.Create(&item).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func NewOrderItemGormRepo(db *gorm.DB) domain.OrderItemRepository {
	return &OrderItemGormRepo{db: db}
}
