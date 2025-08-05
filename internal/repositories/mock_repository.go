package repositories

import (
	"github.com/vimudakorn/internal/domain"
	"gorm.io/gorm"
)

type MockRepository struct {
	db *gorm.DB
}

func NewMockRepository(db *gorm.DB) *MockRepository {
	return &MockRepository{db: db}
}

func (r *MockRepository) InsertMultiple(
	users []domain.User,
	books []domain.Book,
	learnings []domain.LearningSupply,
	products []domain.Product,
) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&users).Error; err != nil {
			return err
		}
		if err := tx.Create(&books).Error; err != nil {
			return err
		}
		if err := tx.Create(&learnings).Error; err != nil {
			return err
		}
		if err := tx.Create(&products).Error; err != nil {
			return err
		}

		// อัพเดต Book.ProductID และ LearningSupply.ProductID ให้ถูกต้อง
		for i := range books {
			books[i].ProductID = products[i+len(learnings)+len(users)].ID // ปรับตามลำดับ
			if err := tx.Save(&books[i]).Error; err != nil {
				return err
			}
		}

		for i := range learnings {
			learnings[i].ProductID = products[i].ID
			if err := tx.Save(&learnings[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
