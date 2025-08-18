package domain

import "gorm.io/gorm"

type UserEduLevel struct {
	gorm.Model
	UserID uint    `gorm:"not null"`                  // FK ไปยัง User
	Levels []Level `gorm:"foreignKey:UserEduLevelID"` // relation
}

type Level struct {
	gorm.Model
	EduLevel       string
	StudentCount   int
	UserEduLevelID uint // FK back to UserEduLevel
}

type UserEduLevelRepository interface {
	CreateWithFixedLevels(userID uint, counts map[string]int) (*UserEduLevel, error)
	GetByUserID(userID uint) (*UserEduLevel, error)
	UpdateStudentCount(userID uint, levelName string, count int) error
	DeleteByUserID(userID uint) error
	UpdateMultipleLevels(userID uint, counts map[string]int) error
}
