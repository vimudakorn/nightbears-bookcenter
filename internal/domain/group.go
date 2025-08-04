package domain

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name        string `gorm:"not null"`
	EduLevel    string `gorm:"not null"`
	Description string
	Products    []GroupProduct
}
