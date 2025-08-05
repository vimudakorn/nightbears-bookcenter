package usecases

import "github.com/vimudakorn/internal/domain"

type CategoryUsecase struct {
	repo domain.CategoryRepository
}

func NewCategoryUsecase(r domain.CategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{repo: r}
}

func (c *CategoryUsecase) Create(category *domain.Category) error {
	return c.repo.Create(category)
}

func (c *CategoryUsecase) IsNameExists(name string) (bool, error) {
	return c.repo.IsNameExists(name)
}

func (c *CategoryUsecase) IsParentExist(id *uint) (bool, error) {
	return c.repo.IsIDExists(id)
}
