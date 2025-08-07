package domain

import (
	groupproductrequest "github.com/vimudakorn/internal/request/group_product_request"
	"gorm.io/gorm"
)

type GroupProduct struct {
	gorm.Model
	GroupID   uint `gorm:"not null"`
	ProductID uint `gorm:"not null"`
	Quantity  int  `gorm:"not null"`
}

type GroupProductRepository interface {
	Create(groupProduct *GroupProduct) error
	CreateMulti(groupProducts []GroupProduct) error
	// GetProductByGroupID(groupID uint) ([]GroupProduct, error)
	GetProductByGroupID(groupID uint) ([]groupproductrequest.GroupProductWithDetail, error)
	CreateWithProduct(tx *gorm.DB, pg *GroupProduct) error
	IsProductInGroupID(groupID uint, productID uint) (bool, error)
	Update(product *GroupProduct) error
	FindByGroupAndProductID(groupID uint, productID uint) (*GroupProduct, error)
	AddOrUpdateMulti(groupID uint, products []GroupProduct) error
}
