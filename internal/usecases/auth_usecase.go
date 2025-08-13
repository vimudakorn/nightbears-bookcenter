package usecases

import (
	"fmt"

	"github.com/vimudakorn/internal/domain"
	"gorm.io/gorm"

	"github.com/vimudakorn/internal/request"
	"github.com/vimudakorn/internal/utils"
)

type AuthUsecase struct {
	userRepo domain.UserRepository
}

func NewAuthUsecase(uRepo domain.UserRepository) *AuthUsecase {
	return &AuthUsecase{userRepo: uRepo}
}

func (u *AuthUsecase) Register(req *request.RegisterRequest) ([]utils.Warning, error) {
	warnings := utils.ValidateRegisterForm(req)
	if len(warnings) > 0 {
		return warnings, nil
	}

	if err := u.IsEmailExist(req.Email); err != nil {
		return nil, fmt.Errorf("Email already exists")
	}

	hashedPassword := utils.HashPassword(req.Password)

	err := u.userRepo.Transaction(func(tx *gorm.DB) error {
		user := &domain.User{
			Email:    req.Email,
			Password: hashedPassword,
			Role:     req.Role,
		}
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		profile := &domain.Profile{
			UserID:    user.ID,
			Name:      req.Name,
			Phone:     req.Phone,
			Address:   req.Address,
			AvatarURL: req.AvatarURL,
		}
		if err := tx.Create(profile).Error; err != nil {
			return err
		}

		cart := &domain.Cart{
			UserID: user.ID,
		}

		if err := tx.Create(cart).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (u *AuthUsecase) IsEmailExist(email string) error {
	existingUser, err := u.userRepo.FindByKey("email", email)
	if err == nil && existingUser != nil {
		return fmt.Errorf("This email already exists")
	}
	return nil
}

func (u *AuthUsecase) Login(email, password string) (*domain.User, error) {
	user, err := u.userRepo.FindByKey("email", email)
	if err != nil {
		return nil, err
	}
	if !utils.CheckPassword(password, user.Password) {
		return nil, fmt.Errorf("password mismatch")
	}
	return user, nil
}

func (u *AuthUsecase) ChangePassword(userID uint, oldPassword, newPassword, confirmNewPassword string) error {
	user, err := u.userRepo.FindByID(userID)
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

	if err := u.userRepo.UpdatePassword(user); err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}
