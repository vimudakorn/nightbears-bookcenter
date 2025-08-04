package domain

import "time"

type Profile struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"unique;not null"`
	User      User   `gorm:"constraint:OnDelete:CASCADE"`
	Name      string `gorm:"not null"`
	Phone     string
	Address   string
	AvatarURL string
	UpdatedAt time.Time
}

type ProfileRepository interface {
	CreateProfile(profile *Profile) error
}
