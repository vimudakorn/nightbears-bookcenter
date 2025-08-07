package repositories

import (
	"errors"

	"github.com/vimudakorn/internal/domain"
	groupproductrequest "github.com/vimudakorn/internal/request/group_product_request"
	"gorm.io/gorm"
)

type GroupProductGormRepo struct {
	db *gorm.DB
}

func (g *GroupProductGormRepo) FindByGroupAndProductID(groupID uint, productID uint) (*domain.GroupProduct, error) {
	var groupProduct domain.GroupProduct
	err := g.db.Where("group_id = ? AND product_id = ?", groupID, productID).First(&groupProduct).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // ไม่เจอ record
		}
		return nil, err // error อื่นๆ
	}
	return &groupProduct, nil
}

// HasProductInGroupID implements domain.GroupProductRepository.
func (g *GroupProductGormRepo) IsProductInGroupID(groupID uint, productID uint) (bool, error) {
	var count int64
	err := g.db.Model(&domain.GroupProduct{}).Find("group_id = ? AND product_id = ?", groupID, productID).Count(&count).Error
	return count > 0, err
}

func (g *GroupProductGormRepo) Update(product *domain.GroupProduct) error {
	return g.db.Model(&domain.GroupProduct{}).
		Where("group_id = ? AND product_id = ?", product.GroupID, product.ProductID).
		Updates(map[string]interface{}{
			"quantity": product.Quantity,
		}).Error
}

func NewGroupProductGormRepo(db *gorm.DB) domain.GroupProductRepository {
	return &GroupProductGormRepo{db: db}
}

// Create implements domain.GroupProductRepository.
func (g *GroupProductGormRepo) Create(groupProduct *domain.GroupProduct) error {
	return g.db.Create(groupProduct).Error
}

func (g *GroupProductGormRepo) CreateMulti(groupProducts []domain.GroupProduct) error {
	return g.db.Transaction(func(tx *gorm.DB) error {
		return tx.Create(&groupProducts).Error
	})
}

// // GetProductByID implements domain.GroupProductRepository.
// func (g *GroupProductGormRepo) GetProductByGroupID(groupID uint) ([]domain.GroupProduct, error) {
// 	var groupProducts []domain.GroupProduct
// 	// err := g.db.
// 	// 	Preload("Product").
// 	// 	Preload("Group").
// 	// 	Where("group_id = ?", groupID).
// 	// 	Find(&groupProducts).Error
// 	err := g.db.Table("group_products").
// 		Joins("JOIN products ON products.id = group_products.product_id").
// 		Joins("JOIN groups ON groups.id = group_products.group_id").
// 		Where("group_products.group_id = ?", groupID).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return groupProducts, nil
// }

func (g *GroupProductGormRepo) GetProductByGroupID(groupID uint) ([]groupproductrequest.GroupProductWithDetail, error) {
	var result []groupproductrequest.GroupProductWithDetail

	err := g.db.Table("group_products").
		Select(`
			group_products.group_id,
			groups.name AS group_name,
			group_products.product_id,
			products.name AS product_name,
			products.price,
			group_products.quantity
		`).
		Joins("JOIN products ON products.id = group_products.product_id").
		Joins("JOIN groups ON groups.id = group_products.group_id").
		Where("group_products.group_id = ?", groupID).
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *GroupProductGormRepo) CreateWithProduct(tx *gorm.DB, pg *domain.GroupProduct) error {
	return tx.Create(pg).Error
}

func (g *GroupProductGormRepo) AddOrUpdateMulti(groupID uint, products []domain.GroupProduct) error {
	return g.db.Transaction(func(tx *gorm.DB) error {
		for _, p := range products {
			var existing domain.GroupProduct
			err := tx.Where("group_id = ? AND product_id = ?", groupID, p.ProductID).First(&existing).Error

			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			if existing.ID != 0 {
				// update quantity
				existing.Quantity += p.Quantity
				if err := tx.Save(&existing).Error; err != nil {
					return err
				}
			} else {
				// create new
				if err := tx.Create(&p).Error; err != nil {
					return err
				}
			}
		}
		return nil
	})
}
