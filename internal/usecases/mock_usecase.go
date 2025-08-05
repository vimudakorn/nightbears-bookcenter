package usecases

import (
	"github.com/vimudakorn/internal/domain"
	"gorm.io/gorm"
)

type MockUseCase struct {
	DB          *gorm.DB
	UserRepo    domain.UserRepository
	ProductRepo domain.ProductRepository
}

func NewMockUseCase(db *gorm.DB, userRepo domain.UserRepository, productRepo domain.ProductRepository) *MockUseCase {
	return &MockUseCase{
		DB:          db,
		UserRepo:    userRepo,
		ProductRepo: productRepo,
	}
}

func (m *MockUseCase) AddNewData() error {
	productsLearning := []domain.Product{
		{ProductCode: 2021580100087, ProductType: "learning", Name: "ซองซิปตาข่าย A4", Price: 20},
		{ProductCode: 8851419014667, ProductType: "learning", Name: "สมุดเส้นเดียว 40 แผ่น", Price: 10},
		{ProductCode: 8858784730205, ProductType: "learning", Name: "ดินน้ำมัน 150 กรัม", Price: 15},
		{ProductCode: 8851907285142, ProductType: "learning", Name: "ดินสอดำ HB", Price: 6},
		{ProductCode: 8851907171063, ProductType: "learning", Name: "ยางลบดินสอ", Price: 6},
		{ProductCode: 8851907075057, ProductType: "learning", Name: "สีเทียน 12 สี มาสเตอร์อาร์ต", Price: 35},
		{ProductCode: 8857121430597, ProductType: "learning", Name: "กระดาษ A4 (แพ็ค 40 แผ่น)", Price: 27},
	}

	learnings := []domain.LearningSupply{
		{Brand: "", Material: "", ProductID: productsLearning[0].ID},
		{Brand: "", Material: "", ProductID: productsLearning[1].ID},
		{Brand: "", Material: "", ProductID: productsLearning[2].ID},
		{Brand: "", Material: "", ProductID: productsLearning[3].ID},
		{Brand: "", Material: "", ProductID: productsLearning[4].ID},
		{Brand: "", Material: "", ProductID: productsLearning[5].ID},
		{Brand: "", Material: "", ProductID: productsLearning[6].ID},
	}

	for i := range productsLearning {
		err := m.ProductRepo.CreateProduct(&productsLearning[i])
		if err != nil {
			return err
		}
		err = m.ProductRepo.CreateLearning(&learnings[i])
		if err != nil {
			return err
		}

	}
	return nil

}
