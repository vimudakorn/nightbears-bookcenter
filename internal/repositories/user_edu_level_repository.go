package repositories

import (
	"github.com/vimudakorn/internal/domain"
	"gorm.io/gorm"
)

type UserEduLevelGormRepo struct {
	db *gorm.DB
}

// IsEduLevelExist implements domain.UserEduLevelRepository.
func (u *UserEduLevelGormRepo) IsEduLevelExist(eduLevelID uint) (bool, error) {
	var count int64
	err := u.db.Model(&domain.UserEduLevel{}).Where("id = ?", eduLevelID).Count(&count).Error
	return count > 0, err
}

// Delete implements domain.UserEduLevelRepository.
func (u *UserEduLevelGormRepo) Delete(userEduLevelID uint) error {
	return u.db.Where("id = ?", userEduLevelID).Delete(&domain.UserEduLevel{}).Error
}

// Update implements domain.UserEduLevelRepository.
func (u *UserEduLevelGormRepo) Update(id uint, update *domain.UserEduLevel) error {
	return u.db.Model(&domain.UserEduLevel{}).
		Where("id = ?", id).
		Updates(update).Error
}

// UpdateMultiEduLevel implements domain.UserEduLevelRepository.
func (u *UserEduLevelGormRepo) UpdateMultiEduLevel(updates []domain.UserEduLevel) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		for _, level := range updates {
			if err := tx.Model(&domain.UserEduLevel{}).
				Where("id = ? AND user_id = ?", level.ID, level.UserID).
				Updates(map[string]interface{}{
					"edu_level":     level.EduLevel,
					"edu_year":      level.EduYear,
					"student_count": level.StudentCount,
				}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// IsEduLevelExist implements domain.UserEduLevelRepository.
func (u *UserEduLevelGormRepo) IsEduLevelNameExist(eduLevel string, eduYear int, userID uint) (bool, error) {
	var count int64
	err := u.db.Model(&domain.UserEduLevel{}).Where("edu_level = ? AND edu_year = ? AND user_id = ?", eduLevel, eduYear, uint(userID)).Count(&count).Error
	return count > 0, err
}

// Create implements domain.UserEduLevelRepository.
func (u *UserEduLevelGormRepo) Create(eduLevel *domain.UserEduLevel) error {
	return u.db.Create(eduLevel).Error
}

func (u *UserEduLevelGormRepo) CreateMultiple(levels []domain.UserEduLevel) error {
	return u.db.Transaction(func(tx *gorm.DB) error {
		for _, level := range levels {
			if err := tx.Create(&level).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
func NewUserEduLevelGormRepo(db *gorm.DB) domain.UserEduLevelRepository {
	return &UserEduLevelGormRepo{db: db}
}
