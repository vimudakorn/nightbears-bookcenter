package repositories

import (
	"fmt"

	"github.com/vimudakorn/internal/domain"
	"gorm.io/gorm"
)

type TagGormRepo struct {
	db *gorm.DB
}

func NewTagGormRepo(db *gorm.DB) domain.TagRepository {
	return &TagGormRepo{db: db}
}

func (r *TagGormRepo) GetTagsByIDs(ids []uint) ([]domain.Tag, error) {
	var tags []domain.Tag
	if len(ids) == 0 {
		return tags, nil
	}
	if err := r.db.Where("id IN ?", ids).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *TagGormRepo) FindAll() ([]domain.Tag, error) {
	var tags []domain.Tag
	err := r.db.
		Preload("Products").Find(&tags).Error
	return tags, err
}

func (r *TagGormRepo) Create(tag *domain.Tag) error {
	return r.db.Create(tag).Error
}

func (r *TagGormRepo) Delete(id uint) error {
	return r.db.Delete(&domain.Tag{}, id).Error
}

func (r *TagGormRepo) Update(tag *domain.Tag) error {
	return r.db.Save(tag).Error
}

func (r *TagGormRepo) GetPagination(page int, limit int, search string, sortBy string, orderBy string) ([]domain.Tag, int64, error) {
	var tags []domain.Tag
	var count int64

	allowedSortBy := map[string]bool{
		"name": true,
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

	query := r.db.Model(&domain.Tag{})

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}

	query.Count(&count)

	err := query.
		Preload("Products").Order(order).Limit(limit).Offset(offset).Find(&tags).Error
	return tags, count, err
}

func (r *TagGormRepo) RenameTag(tagID uint, newName string) error {
	return r.db.Model(&domain.Tag{}).
		Where("id = ?", tagID).
		Update("name", newName).Error
}

func (r *TagGormRepo) IsTagHasUsed(id uint) (bool, error) {
	var count int64

	err := r.db.
		Table("product_tags").
		Where("tag_id = ?", id).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *TagGormRepo) IsNameExists(name string) (bool, error) {
	var count int64
	err := r.db.Model(&domain.Tag{}).Where("name = ?", name).Count(&count).Error

	if err != nil {
		return false, err
	}
	return count > 0, nil
}
