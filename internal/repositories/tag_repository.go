package repositories

import (
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
