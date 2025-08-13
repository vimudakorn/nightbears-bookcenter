package domain

import (
	cartitemresponse "github.com/vimudakorn/internal/responses/cart_item_response"
	"gorm.io/gorm"
)

type CartItem struct {
	gorm.Model
	CartID    uint `gorm:"not null"`
	ProductID *uint
	GroupID   *uint
	Quantity  int `gorm:"default:1;not null"`
}

type CartItemRepository interface {
	AddOrUpdateCartItem(cartID uint, item *CartItem) error
	AddOrUpdateMultiCartItems(cartID uint, items []CartItem) error
	IsProductInCartID(cartID uint, productID uint) (bool, error)
	FindItemByID(cartID uint, productID uint) (*CartItem, error)
	GetItemsByCartID(cartID uint) ([]cartitemresponse.CartItemDetailResponse, error)
	Update(cartItem *CartItem) error
	UpdateItemsInCartID(cartID uint, cartItems []CartItem) error
	Delete(cartID uint, cartItemID uint) error
}
