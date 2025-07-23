package usecases

import (
	"github.com/vimudakorn/internal/domain"
	"github.com/vimudakorn/internal/request"
)

type UserUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(r domain.UserRepository) *UserUsecase {
	return &UserUsecase{repo: r}
}

func (r *UserUsecase) Create(user *domain.User) error {
	return r.repo.Create(user)
}

func (r *UserUsecase) Update(user *domain.User) error {
	return r.repo.Update(user)
}

func (r *UserUsecase) GetAll() ([]domain.User, error) {
	return r.repo.FindAll()
}

func (r *UserUsecase) GetByID(id uint) (*domain.User, error) {
	return r.repo.FindByID(id)
}

func (r *UserUsecase) GetByUsername(name string) (*domain.User, error) {
	return r.repo.FindByName(name)
}

func (r *UserUsecase) Delete(id uint) error {
	return r.repo.Delete(id)
}

func (r *UserUsecase) GetPagination(page int, limit int, search string, sortBy string, sortOrder string) ([]domain.User, int64, error) {
	return r.repo.GetPagination(page, limit, search, sortBy, sortOrder)
}

func (r *UserUsecase) UpdateUserProfile(userID uint, req request.UserProfileDataRequest) error {
	return r.repo.UpdateProfile(userID, req)
}

func (r *UserUsecase) UpdatePassword(user *domain.User) error {
	return r.repo.UpdatePassword(user)
}
