package domain

import "gorm.io/gorm"

type UserEduLevel struct {
	gorm.Model
	UserID       uint   `gorm:"not null"` // foreign key ไปยัง User
	EduLevel     string `gorm:"not null"` // ระดับชั้น เช่น Grade 3
	StudentCount int    `gorm:"not null"` // จำนวนผู้เรียนในระดับนั้น
	EduYear      int    `gorm:"not null"`
}

type UserEduLevelRepository interface {
	Create(eduLevel *UserEduLevel) error
	GetByUserID(userID uint) ([]UserEduLevel, error)
	CreateMultiple(levels []UserEduLevel) error
	IsEduLevelNameExist(eduLevel string, eduYear int, userID uint) (bool, error)
	Update(id uint, update *UserEduLevel) error
	UpdateMultiEduLevel(updates []UserEduLevel) error
	Delete(userEduLevelID uint) error
	IsEduLevelExist(eduLevelID uint) (bool, error)
}
