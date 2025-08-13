package usecases

import (
	"errors"

	"github.com/vimudakorn/internal/domain"
)

type UserEduLevelUsecase struct {
	eduLevelRepo domain.UserEduLevelRepository
}

func NewUserEduLevelUsecase(elRepo domain.UserEduLevelRepository) *UserEduLevelUsecase {
	return &UserEduLevelUsecase{eduLevelRepo: elRepo}
}

func (el *UserEduLevelUsecase) Create(eduLevel *domain.UserEduLevel) error {
	return el.eduLevelRepo.Create(eduLevel)
}

func (el *UserEduLevelUsecase) CreateMultiple(levels []domain.UserEduLevel) error {
	if len(levels) == 0 {
		return nil
	}
	return el.eduLevelRepo.CreateMultiple(levels)
}

func (el *UserEduLevelUsecase) IsEduLevelNameExist(eduLevel string, eduYear int, userID uint) (bool, error) {
	return el.eduLevelRepo.IsEduLevelNameExist(eduLevel, eduYear, userID)
}

func (el *UserEduLevelUsecase) IsEduLevelExist(eduLevelID uint) (bool, error) {
	return el.eduLevelRepo.IsEduLevelExist(eduLevelID)
}

func (el *UserEduLevelUsecase) UpdateMultiple(userID uint, levels []domain.UserEduLevel) error {
	if len(levels) == 0 {
		return errors.New("no education levels to update")
	}

	for i := range levels {
		levels[i].UserID = userID
	}

	return el.eduLevelRepo.UpdateMultiEduLevel(levels)
}

func (el *UserEduLevelUsecase) UpdateByID(id uint, update *domain.UserEduLevel) error {
	return el.eduLevelRepo.Update(id, update)
}

func (el *UserEduLevelUsecase) Delete(userEduLevelID uint) error {
	return el.eduLevelRepo.Delete(userEduLevelID)
}
