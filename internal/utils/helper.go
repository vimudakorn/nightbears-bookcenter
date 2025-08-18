package utils

import (
	"math/rand"
	"time"

	"github.com/vimudakorn/internal/domain"
	"gorm.io/gorm"
)

func GenerateUniqueProductCode(db *gorm.DB) (int, error) {
	rand.Seed(time.Now().UnixNano())
	for {
		code := rand.Intn(900000) + 100000
		var count int64
		if err := db.Model(&domain.Product{}).Where("product_code = ?", code).Count(&count).Error; err != nil {
			return 0, err
		}
		if count == 0 {
			return code, nil
		}
	}
}
