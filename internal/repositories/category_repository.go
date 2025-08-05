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
	return c.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&domain.Category{}).
			Where("parent_id = ?", id).
			Update("parent_id", nil).Error; err != nil {
			return err
		}

		result := tx.Delete(&domain.Category{}, id)
		if result.RowsAffected == 0 {
			return errors.New("category not found")
		}

		return result.Error
	})
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

func (r *CategoryGormRepo) HasChildren(id uint) (bool, error) {
	var count int64
	err := r.db.Model(&domain.Category{}).Where("parent_id = ?", id).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *CategoryGormRepo) HasProducts(categoryID uint) (bool, error) {
	var count int64
	err := r.db.Model(&domain.Product{}).Where("category_id = ?", categoryID).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
