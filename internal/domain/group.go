package domain

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name        string `gorm:"not null"`
	EduLevel    string `gorm:"not null"`
	Description string
	SalePrice   float64
	Products    []GroupProduct `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE;"`
}

type GroupRepository interface {
	GetPagination(page int, limit int, search string, sortBy string, orderBy string) ([]Group, int64, error)
	Create(group *Group) error
	Update(group *Group) error
	Delete(id uint) error
	IsNameAndEduExist(name string, level string) (bool, error)
	FindByID(id uint) (*Group, error)
	CreateWithProduct(tx *gorm.DB, group *Group) error
	GetDB() *gorm.DB
	IsGroupIDExists(groupID uint) (bool, error)
}
