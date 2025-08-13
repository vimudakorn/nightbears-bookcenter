package repositories

import (
	"fmt"

	"github.com/vimudakorn/internal/domain"
	"github.com/vimudakorn/internal/request"
	"gorm.io/gorm"
)

type UserGormRepo struct {
	db *gorm.DB
}

func NewUserGormRepo(db *gorm.DB) domain.UserRepository {
	return &UserGormRepo{db: db}
}

func (u *UserGormRepo) FindByKey(key string, value string) (*domain.User, error) {
	user := domain.User{}

	// Avoid SQL injection by whitelisting allowed columns
	allowedKeys := map[string]bool{
		"id": true, "email": true, "phone": true, "username": true,
	}

	if !allowedKeys[key] {
		return nil, fmt.Errorf("invalid query key: %s", key)
	}

	err := u.db.Where(fmt.Sprintf("%s = ?", key), value).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserGormRepo) Create(user *domain.User) error {
	return u.db.Create(user).Error
}

func (u *UserGormRepo) Delete(id uint) error {
	return u.db.Delete(&domain.User{}, id).Error
}

func (u *UserGormRepo) FindAll() ([]domain.User, error) {
	users := []domain.User{}
	err := u.db.Find(&users).Error
	return users, err
}

func (u *UserGormRepo) FindByID(id uint) (*domain.User, error) {
	user := domain.User{}
	// err := u.db.Preload("Cart").Preload("Order").Preload("Book").First(&user, id).Error
	err := u.db.Preload("Profile").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserGormRepo) FindByName(name string) (*domain.User, error) {
	user := domain.User{}
	err := u.db.Where("name = ?", name).First(&user).Error
	return &user, err
}

func (u *UserGormRepo) GetPagination(page int, limit int, search string, sortBy string, sortOrder string) ([]domain.User, int64, error) {
	var users []domain.User
	var count int64

	allowedSortBy := map[string]bool{
		"name":  true,
		"email": true,
		"role":  true,
		"id":    true,
	}
	allowedSortOrder := map[string]bool{
		"asc":  true,
		"desc": true,
	}

	if !allowedSortBy[sortBy] {
		sortBy = "id"
	}
	if !allowedSortOrder[sortOrder] {
		sortOrder = "asc"
	}

	offset := (page - 1) * limit
	order := fmt.Sprintf("%s %s", sortBy, sortOrder)

	query := u.db.Model(&domain.User{})

	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
		// query = query.Where("email ILIKE ?", "%"+search+"%")
	}

	query.Count(&count)

	err := query.Preload("Profile").
		Preload("Cart").
		Preload("Orders").Order(order).Limit(limit).Offset(offset).Find(&users).Error
	return users, count, err
}

func (u *UserGormRepo) Update(user *domain.User) error {
	return u.db.Save(user).Error
}

func (r *UserGormRepo) UpdateProfile(userID uint, req request.UserProfileDataRequest) error {
	updates := map[string]interface{}{}
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Phone != "" {
		updates["phone"] = req.Phone
	}

	return r.db.Model(&domain.Profile{}).Where("user_id = ?", userID).Updates(updates).Error
}

func (u *UserGormRepo) UpdatePassword(user *domain.User) error {
	return u.db.Model(user).Update("password", user.Password).Error
}

func (r *UserGormRepo) CreateProfile(profile *domain.Profile) error {
	return r.db.Create(profile).Error
}

func (r *UserGormRepo) Transaction(fc func(tx *gorm.DB) error) error {
	return r.db.Transaction(fc)
}
