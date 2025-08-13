package usecases

import (
	"github.com/vimudakorn/internal/domain"
	cartitemresponse "github.com/vimudakorn/internal/responses/cart_item_response"
)

type CartItemUsecase struct {
	cartItemRepo domain.CartItemRepository
	cartRepo     domain.CartRepository
	groupRepo    domain.GroupRepository
	productRepo  domain.ProductRepository
}

func NewCartItemUsecase(ciRepo domain.CartItemRepository, cRepo domain.CartRepository, gRepo domain.GroupRepository, pRepo domain.ProductRepository) *CartItemUsecase {
	return &CartItemUsecase{cartItemRepo: ciRepo, cartRepo: cRepo, groupRepo: gRepo, productRepo: pRepo}
}

func (ci *CartItemUsecase) GetItemsByCartID(cartID uint) ([]cartitemresponse.CartItemDetailResponse, error) {
	return ci.cartItemRepo.GetItemsByCartID(cartID)
}

func (ci *CartItemUsecase) AddOrUpdateMultiCartItems(cartID uint, items []domain.CartItem) error {
	return ci.cartItemRepo.AddOrUpdateMultiCartItems(cartID, items)
}

func (ci *CartItemUsecase) AddOrUpdateCartItem(cartID uint, item *domain.CartItem) error {
	return ci.cartItemRepo.AddOrUpdateCartItem(cartID, item)
}

func (ci *CartItemUsecase) GetCartByUserID(userID uint) (*domain.Cart, error) {
	return ci.cartRepo.GetByUserID(userID)
}

func (ci *CartItemUsecase) IsGroupIDExist(groupID uint) (bool, error) {
	return ci.groupRepo.IsGroupIDExists(groupID)
}

func (ci *CartItemUsecase) IsProductIDExists(productID uint) (bool, error) {
	return ci.productRepo.IsProductIDExists(productID)
}

func (ci *CartItemUsecase) Update(cartItem *domain.CartItem) error {
	return ci.cartItemRepo.Update(cartItem)
}

func (ci *CartItemUsecase) UpdateItemInCartID(cartID uint, cartItems []domain.CartItem) error {
	return ci.cartItemRepo.UpdateItemsInCartID(cartID, cartItems)
}

func (ci *CartItemUsecase) Delete(cartID uint, cartItemID uint) error {
	return ci.cartItemRepo.Delete(cartID, cartItemID)
}
