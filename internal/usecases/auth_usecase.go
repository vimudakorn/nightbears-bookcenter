package usecases

import (
	"fmt"

	"github.com/vimudakorn/internal/domain"

	"github.com/vimudakorn/internal/request"
	"github.com/vimudakorn/internal/utils"
)

type AuthUsecase struct {
	repo domain.UserRepository
}

func NewAuthUsecase(r domain.UserRepository) *AuthUsecase {
	return &AuthUsecase{repo: r}
}

func (u *AuthUsecase) Register(body *request.RegisterRequest) error {
	// existingUser, err := u.repo.FindByKey("email", body.Email)
	// if err == nil && existingUser != nil {
	// 	return fmt.Errorf("This email already exists")
	// }
	if err := u.IsEmailExist(body.Email); err != nil {
		return err
	}
	if body.Password != body.ConfirmPassword {
		return fmt.Errorf("passwords do not match")
	}
	hashedPassword := utils.HashPassword(body.Password)
	user := &domain.User{Name: body.Name, Email: body.Email, Password: hashedPassword, Role: body.Role, Phone: body.Phone}
	return u.repo.Create(user)
}

func (u *AuthUsecase) IsEmailExist(email string) error {
	existingUser, err := u.repo.FindByKey("email", email)
	if err == nil && existingUser != nil {
		return fmt.Errorf("This email already exists")
	}
	return nil
}

func (u *AuthUsecase) Login(email, password string) (*domain.User, error) {
	user, err := u.repo.FindByKey("email", email)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPassword(password, user.Password) {
		return nil, fmt.Errorf("password mismatch")
	}
	return user, nil
}

func (u *AuthUsecase) ChangePassword(userID uint, oldPassword, newPassword, confirmNewPassword string) error {
	user, err := u.repo.FindByID(userID)
	if err != nil {
		return fmt.Errorf("user not found")
	}

	if newPassword != confirmNewPassword {
		return fmt.Errorf("password and confirm password mismatch")
	}

	if !utils.CheckPassword(oldPassword, user.Password) {
		return fmt.Errorf("old password is incorrect")
	}

	user.Password = utils.HashPassword(newPassword)

	if err := u.repo.UpdatePassword(user); err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}
