package repositories

import (
	"fmt"

	"github.com/vimudakorn/internal/domain"
	"gorm.io/gorm"
)

type GroupGormRepo struct {
	db *gorm.DB
}

// IsGroupIDExists implements domain.GroupRepository.
func (g *GroupGormRepo) IsGroupIDExists(groupID uint) (bool, error) {
	var count int64
	err := g.db.Model(&domain.Group{}).Where("id = ?", groupID).Count(&count).Error
	return count > 0, err
}

func NewGroupGormRepo(db *gorm.DB) domain.GroupRepository {
	return &GroupGormRepo{db: db}
}

// Create implements domain.GroupRepository.
func (g *GroupGormRepo) Create(group *domain.Group) error {
	return g.db.Create(group).Error
}

// Delete implements domain.GroupRepository.
func (g *GroupGormRepo) Delete(id uint) error {
	return g.db.Delete(&domain.Group{}, id).Error
}

// GetPagination implements domain.GroupRepository.
func (g *GroupGormRepo) GetPagination(page int, limit int, search string, sortBy string, orderBy string) ([]domain.Group, int64, error) {
	var groups []domain.Group
	var count int64

	allowedSortBy := map[string]bool{
		"name":      true,
		"edu_level": true,
	}
	allowedOrderBy := map[string]bool{
		"asc":  true,
		"desc": true,
	}

	if !allowedSortBy[sortBy] {
		sortBy = "id"
	}
	if !allowedOrderBy[orderBy] {
		orderBy = "asc"
	}

	offset := (page - 1) * limit
	order := fmt.Sprintf("%s %s", sortBy, orderBy)

	query := g.db.Model(&domain.Group{})

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	// นับจำนวนทั้งหมดก่อน
	query.Count(&count)

	err := query.
		Preload("Products", func(db *gorm.DB) *gorm.DB {
			return db.
				Joins("JOIN products ON products.id = group_products.product_id").
				Where("products.deleted_at IS NULL")
		}).
		Order(order).
		Limit(limit).
		Offset(offset).
		Find(&groups).Error

	return groups, count, err
}

// Update implements domain.GroupRepository.
func (g *GroupGormRepo) Update(group *domain.Group) error {
	return g.db.Save(group).Error
}

// IsNameAndEduExist implements domain.GroupRepository.
func (g *GroupGormRepo) IsNameAndEduExist(name string, level string) (bool, error) {
	var count int64

	err := g.db.Model(&domain.Group{}).Where("name = ? AND edu_level = ?", name, level).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (g *GroupGormRepo) FindByID(id uint) (*domain.Group, error) {
	var group domain.Group
	err := g.db.Preload("Products").First(&group, id).Error
	return &group, err
}

func (r *GroupGormRepo) CreateWithProduct(tx *gorm.DB, group *domain.Group) error {
	return tx.Create(group).Error
}

func (r *GroupGormRepo) GetDB() *gorm.DB {
	return r.db
}
