package usecases

import (
	"github.com/vimudakorn/internal/domain"
)

type UserEduLevelUsecase struct {
	eduLevelRepo domain.UserEduLevelRepository
}

func NewUserEduLevelUsecase(elRepo domain.UserEduLevelRepository) *UserEduLevelUsecase {
	return &UserEduLevelUsecase{eduLevelRepo: elRepo}
}

func (u *UserEduLevelUsecase) CreateWithFixedLevels(userID uint, counts map[string]int) (*domain.UserEduLevel, error) {
	return u.eduLevelRepo.CreateWithFixedLevels(userID, counts)
}

func (u *UserEduLevelUsecase) UpdateMultipleLevels(userID uint, counts map[string]int) error {
	return u.eduLevelRepo.UpdateMultipleLevels(userID, counts)
}
func (u *UserEduLevelUsecase) GetByUserID(userID uint) (*domain.UserEduLevel, error) {
	return u.eduLevelRepo.GetByUserID(userID)
}

// อัปเดตจำนวน นร. ในแต่ละชั้น
func (u *UserEduLevelUsecase) UpdateStudentCount(userID uint, levelName string, count int) error {
	return u.eduLevelRepo.UpdateStudentCount(userID, levelName, count)
}

// ลบ UserEduLevel (รวมทั้ง levels ด้วย)
func (u *UserEduLevelUsecase) DeleteByUserID(userID uint) error {
	return u.eduLevelRepo.DeleteByUserID(userID)
}
