package domain

type ProductCategory struct {
	ProductID  uint `gorm:"primaryKey"`
	CategoryID uint `gorm:"primaryKey"`
}
