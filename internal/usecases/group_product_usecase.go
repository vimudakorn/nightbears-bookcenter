package usecases

import "github.com/vimudakorn/internal/domain"

type GroupProductUsecase struct {
	repo domain.GroupProductRepository
}

func NewGroupProductUsecase(r domain.GroupProductRepository) *GroupProductUsecase {
	return &GroupProductUsecase{repo: r}
}
