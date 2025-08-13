package usecases

import "github.com/vimudakorn/internal/domain"

type CartUsecase struct {
	cartRepo domain.CartRepository
}

func NewCartUsecase(cRepo domain.CartRepository) *CartUsecase {
	return &CartUsecase{cartRepo: cRepo}
}

func (c *CartUsecase) Create(cart *domain.Cart) error {
	return c.cartRepo.Create(cart)
}

func (c *CartUsecase) GetCartByUserID(userID uint) (*domain.Cart, error) {
	return c.cartRepo.GetByUserID(userID)
}
