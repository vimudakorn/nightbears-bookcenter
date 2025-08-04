package domain

import (
	"github.com/vimudakorn/internal/request"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
	Role     string `gorm:"not null"`
	Profile  *Profile
	Cart     Cart    `gorm:"foreignKey:UserID"`
	Orders   []Order `gorm:"foreignKey:UserID"`
}

type UserRepository interface {
	Transaction(fc func(tx *gorm.DB) error) error
	Create(user *User) error
	CreateProfile(Profile *Profile) error
	FindAll() ([]User, error)
	FindByName(name string) (*User, error)
	FindByKey(key string, value string) (*User, error)
	// SearchUsers(name string, limit, offset int) ([]User, int64, error)
	GetPagination(page int, limit int, search string, sortBy string, sortOrder string) ([]User, int64, error)
	FindByID(id uint) (*User, error)
	Update(user *User) error
	UpdateProfile(userID uint, req request.UserProfileDataRequest) error
	UpdatePassword(user *User) error
	Delete(id uint) error
}
