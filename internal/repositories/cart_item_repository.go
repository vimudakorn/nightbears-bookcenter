package repositories

import (
	"errors"
	"fmt"

	"github.com/vimudakorn/internal/domain"
	cartitemresponse "github.com/vimudakorn/internal/responses/cart_item_response"
	"gorm.io/gorm"
)

type CartItemGormRepo struct {
	db *gorm.DB
}

func (c *CartItemGormRepo) Delete(cartID uint, cartItemID uint) error {
	var cart domain.Cart
	if err := c.db.First(&cart, cartID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("cart with ID %d does not exist", cartID)
		}
		return err
	}

	var cartItem domain.CartItem
	if err := c.db.Where("id = ? AND cart_id = ?", cartItemID, cartID).First(&cartItem).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("cart item with ID %d not found in cart %d", cartItemID, cartID)
		}
		return err
	}

	if err := c.db.Delete(&cartItem).Error; err != nil {
		return err
	}

	return nil
}

func (c *CartItemGormRepo) Update(cartItem *domain.CartItem) error {
	query := c.db.Model(&domain.CartItem{}).Where("cart_id = ?", cartItem.CartID)

	if cartItem.ProductID != nil {
		query = query.Where("product_id = ?", *cartItem.ProductID)
	} else {
		query = query.Where("product_id IS NULL")
	}

	if cartItem.GroupID != nil {
		query = query.Where("group_id = ?", *cartItem.GroupID)
	} else {
		query = query.Where("group_id IS NULL")
	}

	return query.Updates(map[string]interface{}{
		"quantity": cartItem.Quantity,
	}).Error
}

// UpdateItemsInCartID implements domain.CartItemRepository.
func (c *CartItemGormRepo) UpdateItemsInCartID(cartID uint, cartItems []domain.CartItem) error {
	return c.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range cartItems {
			query := tx.Model(&domain.CartItem{}).Where("cart_id = ?", cartID)

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

			var existing domain.CartItem
			err := query.First(&existing).Error

			if errors.Is(err, gorm.ErrRecordNotFound) {
				// ถ้าไม่เจอ ให้ข้าม
				continue
			}
			if err != nil {
				return err
			}

			// อัปเดต quantity
			existing.Quantity = item.Quantity
			if err := tx.Save(&existing).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// GetItemsByCartID implements domain.CartItemRepository.
func (c *CartItemGormRepo) GetItemsByCartID(cartID uint) ([]cartitemresponse.CartItemDetailResponse, error) {
	var result []cartitemresponse.CartItemDetailResponse

	err := c.db.Table("cart_items ci").
		Select(`
			ci.cart_id,
			u.id AS user_id,
			u.email AS user_email,
			ci.product_id,
			p.product_type,
			p.name AS product_name,
			p.price AS product_price,
			ci.group_id,
			g.name AS group_name,
			g.sale_price AS group_price,
			ci.quantity,

			-- Book details
			b.subject AS book_subject,
			b.learning_area AS book_learning_area,
			b.grade AS book_grade,
			b.publisher AS book_publisher,
			b.author AS book_author,
			b.isbn AS book_isbn,

			-- Learning supply details
			ls.brand AS learning_brand,
			ls.material AS learning_material,

			-- Office supply details
			os.color AS office_color,
			os.size AS office_size
		`).
		Joins("JOIN carts c ON c.id = ci.cart_id AND c.deleted_at IS NULL").
		Joins("JOIN users u ON u.id = c.user_id").
		Joins("LEFT JOIN products p ON p.id = ci.product_id AND p.deleted_at IS NULL").
		Joins("LEFT JOIN groups g ON g.id = ci.group_id AND g.deleted_at IS NULL").
		Joins("LEFT JOIN books b ON b.product_id = p.id").
		Joins("LEFT JOIN learning_supplies ls ON ls.product_id = p.id").
		Joins("LEFT JOIN office_supplies os ON os.product_id = p.id").
		Where("ci.cart_id = ? AND ci.deleted_at IS NULL", cartID).
		Scan(&result).Error

	return result, err
}

// func (c *CartItemGormRepo) GetItemsByCartID(cartID uint) ([]cartitemresponse.CartItemDetailResponse, error) {
// 	var result []cartitemresponse.CartItemDetailResponse

// 	// First: Get main cart item with product/group basic info
// 	err := c.db.Table("cart_items ci").
// 		Select(`
// 			ci.cart_id,
// 			u.id AS user_id,
// 			u.email AS user_email,
// 			ci.product_id,
// 			p.name AS product_name,
// 			p.price AS product_price,
// 			ci.group_id,
// 			g.name AS group_name,
// 			g.sale_price AS group_price,
// 			ci.quantity
// 		`).
// 		Joins("JOIN carts c ON c.id = ci.cart_id AND c.deleted_at IS NULL").
// 		Joins("JOIN users u ON u.id = c.user_id").
// 		Joins("LEFT JOIN products p ON p.id = ci.product_id AND p.deleted_at IS NULL").
// 		Joins("LEFT JOIN groups g ON g.id = ci.group_id AND g.deleted_at IS NULL").
// 		Where("ci.cart_id = ? AND ci.deleted_at IS NULL", cartID).
// 		Scan(&result).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	// Second: For each group, fetch group items and product details
// 	for i, item := range result {
// 		if item.GroupID != nil {
// 			var groupProducts []cartitemresponse.GroupProductDetail
// 			err := c.db.Table("group_products gp").
// 				Select(`
// 					gp.product_id,
// 					p.name AS product_name,
// 					p.price AS product_price,
// 					p.product_type,
// 					b.subject, b.learning_area, b.grade, b.publisher, b.author, b.isbn,
// 					ls.brand, ls.material,
// 					os.color, os.size
// 				`).
// 				Joins("JOIN products p ON p.id = gp.product_id").
// 				Joins("LEFT JOIN books b ON b.product_id = p.id").
// 				Joins("LEFT JOIN learning_supplies ls ON ls.product_id = p.id").
// 				Joins("LEFT JOIN office_supplies os ON os.product_id = p.id").
// 				Where("gp.group_id = ? AND p.deleted_at IS NULL", *item.GroupID).
// 				Scan(&groupProducts).Error

// 			if err != nil {
// 				return nil, err
// 			}

// 			// Attach group products back
// 			result[i].GroupDetail = cartitemresponse.GroupDetail{
// 				GroupID:    *item.GroupID,
// 				GroupName:  item.GroupName,
// 				GroupPrice: item.GroupPrice,
// 				Products:   groupProducts,
// 			}
// 		}
// 	}

// 	return result, nil
// }

// FindItemByID implements domain.CartItemRepository.
func (c *CartItemGormRepo) FindItemByID(cartID uint, productID uint) (*domain.CartItem, error) {
	var cartItem domain.CartItem
	err := c.db.Model(&domain.CartItem{}).Where("cart_id = ? AND product_id = ?", cartID, productID).First(&cartItem).Error
	return &cartItem, err
}

// IsProductInCartID implements domain.CartItemRepository.
func (c *CartItemGormRepo) IsProductInCartID(cartID uint, productID uint) (bool, error) {
	var count int64
	err := c.db.Model(&domain.CartItem{}).Where("cart_id = ? AND product_id = ?", cartID, productID).Count(&count).Error
	return count > 0, err
}

// AddOrUpdateCartItem implements domain.CartItemRepository.
func (c *CartItemGormRepo) AddOrUpdateCartItem(cartID uint, item *domain.CartItem) error {
	// Check if cart exists
	var cart domain.Cart
	if err := c.db.First(&cart, cartID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("cart with ID %d does not exist", cartID)
		}
		return err
	}

	// Validate product_id or group_id
	if (item.ProductID != nil && item.GroupID != nil) || (item.ProductID == nil && item.GroupID == nil) {
		return errors.New("either product_id or group_id must be set, but not both")
	}

	var existing domain.CartItem
	var err error

	if item.ProductID != nil {
		err = c.db.Where("cart_id = ? AND product_id = ?", cartID, *item.ProductID).
			First(&existing).Error
	} else {
		err = c.db.Where("cart_id = ? AND group_id = ?", cartID, *item.GroupID).
			First(&existing).Error
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if existing.ID != 0 {
		existing.Quantity += item.Quantity
		return c.db.Save(&existing).Error
	}

	item.CartID = cartID
	return c.db.Create(item).Error
}

func (r *CartItemGormRepo) AddOrUpdateMultiCartItems(cartID uint, items []domain.CartItem) error {
	var cart domain.Cart
	if err := r.db.First(&cart, cartID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("cart with ID %d does not exist", cartID)
		}
		return err
	}

	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, item := range items {
			if (item.ProductID != nil && item.GroupID != nil) || (item.ProductID == nil && item.GroupID == nil) {
				return errors.New("either product_id or group_id must be set, but not both")
			}

			var existing domain.CartItem
			var err error

			if item.ProductID != nil {
				err = tx.Where("cart_id = ? AND product_id = ?", cartID, *item.ProductID).First(&existing).Error
			} else {
				err = tx.Where("cart_id = ? AND group_id = ?", cartID, *item.GroupID).First(&existing).Error
			}

			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			if existing.ID != 0 {
				// update quantity
				existing.Quantity += item.Quantity
				if err := tx.Save(&existing).Error; err != nil {
					return err
				}
			} else {
				item.CartID = cartID
				if err := tx.Create(&item).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func NewCartItemGormRepo(db *gorm.DB) domain.CartItemRepository {
	return &CartItemGormRepo{db: db}
}
