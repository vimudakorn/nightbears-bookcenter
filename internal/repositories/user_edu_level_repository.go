package repositories

import (
	"github.com/vimudakorn/internal/domain"
	"gorm.io/gorm"
)

type UserEduLevelGormRepo struct {
	db *gorm.DB
}

// Fix ค่า EduLevel ไว้ล่วงหน้า
var fixedLevels = []string{
	"อนุบาลศึกษาปีที่ 1",
	"อนุบาลศึกษาปีที่ 2",
	"อนุบาลศึกษาปีที่ 3",
	"ประถมศึกษาปีที่ 1",
	"ประถมศึกษาปีที่ 2",
	"ประถมศึกษาปีที่ 3",
	"ประถมศึกษาปีที่ 4",
	"ประถมศึกษาปีที่ 5",
	"ประถมศึกษาปีที่ 6",
	"มัธยมศึกษาปีที่ 1",
	"มัธยมศึกษาปีที่ 2",
	"มัธยมศึกษาปีที่ 3",
	"มัธยมศึกษาปีที่ 4",
	"มัธยมศึกษาปีที่ 5",
	"มัธยมศึกษาปีที่ 6",
}

// สร้าง UserEduLevel พร้อม Levels ตาม fixedLevels
// input map[levelName]studentCount (ถ้าไม่ส่ง จะเป็น 0)
func (u *UserEduLevelGormRepo) CreateWithFixedLevels(userID uint, counts map[string]int) (*domain.UserEduLevel, error) {
	edu := domain.UserEduLevel{UserID: userID}

	for _, lvl := range fixedLevels {
		count := 0
		if val, ok := counts[lvl]; ok {
			count = val
		}
		edu.Levels = append(edu.Levels, domain.Level{
			EduLevel:     lvl,
			StudentCount: count,
		})
	}

	err := u.db.Create(&edu).Error
	return &edu, err
}

// Update หลายระดับพร้อมกัน
func (u *UserEduLevelGormRepo) UpdateMultipleLevels(userID uint, counts map[string]int) error {
	for lvl, count := range counts {
		if err := u.db.Model(&domain.Level{}).
			Where("edu_level = ? AND user_edu_level_id = (?)",
				lvl,
				u.db.Table("user_edu_levels").Select("id").Where("user_id = ?", userID),
			).
			Update("student_count", count).Error; err != nil {
			return err
		}
	}
	return nil
}

// ✅ Get UserEduLevel + Levels
func (u *UserEduLevelGormRepo) GetByUserID(userID uint) (*domain.UserEduLevel, error) {
	var edu domain.UserEduLevel
	err := u.db.Preload("Levels").
		Where("user_id = ?", userID).
		First(&edu).Error
	return &edu, err
}

// ✅ Update จำนวนนักเรียนในแต่ละชั้น
func (u *UserEduLevelGormRepo) UpdateStudentCount(userID uint, levelName string, count int) error {
	return u.db.Model(&domain.Level{}).
		Where("edu_level = ? AND user_edu_level_id = ?", levelName, userID).
		Update("student_count", count).Error
}

// DeleteByUserID ลบ UserEduLevel + Levels ของ user
func (u *UserEduLevelGormRepo) DeleteByUserID(userID uint) error {
	// ลบ levels ก่อน (foreign key)
	if err := u.db.Where("user_edu_level_id IN (?)",
		u.db.Table("user_edu_levels").Select("id").Where("user_id = ?", userID),
	).Delete(&domain.Level{}).Error; err != nil {
		return err
	}

	// ลบ UserEduLevel
	if err := u.db.Where("user_id = ?", userID).Delete(&domain.UserEduLevel{}).Error; err != nil {
		return err
	}

	return nil
}

func NewUserEduLevelGormRepo(db *gorm.DB) domain.UserEduLevelRepository {
	return &UserEduLevelGormRepo{db: db}
}
