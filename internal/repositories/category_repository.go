package repositories

import (
	"errors"

	"github.com/vimudakorn/internal/domain"
	"gorm.io/gorm"
)

type CategoryGormRepo struct {
	db *gorm.DB
}

func NewCategoryGormRepo(db *gorm.DB) domain.CategoryRepository {
	return &CategoryGormRepo{db: db}
}

func (c *CategoryGormRepo) IsIDExists(id *uint) (bool, error) {
	if id == nil {
		return true, nil
	}

	var count int64
	err := c.db.Model(&domain.Category{}).Where("id = ?", *id).Count(&count).Error
	return count > 0, err
}

func (c *CategoryGormRepo) Create(category *domain.Category) error {
	return c.db.Create(category).Error
}

func (c *CategoryGormRepo) Delete(id uint) error {
	result := c.db.Delete(&domain.Category{}, id)
	if result.RowsAffected == 0 {
		return errors.New("Category not found")
	}
	return result.Error
}

func (c *CategoryGormRepo) FindAll() ([]domain.Category, error) {
	var categories []domain.Category
	err := c.db.Preload("Children").Find(&categories).Error
	return categories, err
}

func (c *CategoryGormRepo) FindByID(id uint) (*domain.Category, error) {
	var category domain.Category
	err := c.db.Preload("Children").First(&category, id).Error
	return &category, err
}

func (c *CategoryGormRepo) FindChildren(parentID uint) ([]domain.Category, error) {
	var children []domain.Category
	err := c.db.Where("parent_id = ?", parentID).Find(&children).Error
	return children, err
}

func (c *CategoryGormRepo) FindRootCategories() ([]domain.Category, error) {
	var roots []domain.Category
	err := c.db.Preload("Children").Where("parent_id IS NULL").Find(&roots).Error
	return roots, err
}

func (c *CategoryGormRepo) IsNameExists(name string) (bool, error) {
	var count int64
	err := c.db.Model(&domain.Category{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}

func (c *CategoryGormRepo) Update(category *domain.Category) error {
	return c.db.Save(category).Error
}

func (c *CategoryGormRepo) FindAllDescendants(id uint) ([]domain.Category, error) {
	var allDescendants []domain.Category

	err := c.FindChildrenRecursive(id, &allDescendants)
	if err != nil {
		return nil, err
	}

	return allDescendants, nil
}

func (c *CategoryGormRepo) FindChildrenRecursive(parentID uint, collected *[]domain.Category) error {
	var children []domain.Category
	if err := c.db.Where("parent_id = ?", parentID).Find(&children).Error; err != nil {
		return err
	}

	for _, child := range children {
		*collected = append(*collected, child)
		if err := c.FindChildrenRecursive(child.ID, collected); err != nil {
			return err
		}
	}

	return nil
}
