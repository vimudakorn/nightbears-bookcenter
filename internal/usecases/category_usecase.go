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

func (c *CategoryUsecase) GetAll() ([]domain.Category, error) {
	return c.repo.FindAll()
}

func (c *CategoryUsecase) FindByID(id uint) (*domain.Category, error) {
	return c.repo.FindByID(id)
}

func (c *CategoryUsecase) Update(category *domain.Category) error {
	return c.repo.Update(category)
}

func (c *CategoryUsecase) Delete(id uint) error {
	return c.repo.Delete(id)
}

func (c *CategoryUsecase) HasChildren(id uint) (bool, error) {
	return c.repo.HasChildren(id)
}

func (c *CategoryUsecase) HasProducts(categoryID uint) (bool, error) {
	return c.repo.HasProducts(categoryID)
}

func (c *CategoryUsecase) IsIDExists(id *uint) (bool, error) {
	return c.repo.IsIDExists(id)
}
