package usecases

import (
	"github.com/vimudakorn/internal/domain"
	"github.com/vimudakorn/internal/repositories"
	"github.com/vimudakorn/internal/utils"
	"gorm.io/gorm"
)

// import (
// 	"github.com/vimudakorn/internal/domain"
// 	"github.com/vimudakorn/internal/utils"
// 	"gorm.io/gorm"
// )

// type InsertMultipleUseCase struct {
// 	DB          *gorm.DB
// 	UserRepo    domain.UserRepository
// 	ProductRepo domain.ProductRepository
// }

// func NewInsertMultipleUseCase(db *gorm.DB, userRepo domain.UserRepository, productRepo domain.ProductRepository) *InsertMultipleUseCase {
// 	return &InsertMultipleUseCase{
// 		DB:          db,
// 		UserRepo:    userRepo,
// 		ProductRepo: productRepo,
// 	}
// }

// func (uc *InsertMultipleUseCase) Execute() error {
// 	return uc.DB.Transaction(func(tx *gorm.DB) error {
// 		users := []domain.User{
// 			{
// 				Email:    "nut1@mail.com",
// 				Password: utils.HashPassword("12341234"),
// 				Role:     "ADMIN",
// 			},
// 			{
// 				Email: "nut2@mail.com",
// 				Password: utils.HashPassword("12341234"),
// 				Role: "USER",
// 			},
// 		}

// 		products := []domain.Product{
// 			{
// 				ProductCode: 2021580100087,
// 				ProductType: "learning",
// 				Name: "ซองซิปตาข่าย A4",
// 				Price: 20,
// 				LearningSupplyID: 1,
// 			},
// 			{
// 				ProductCode: 8851419014667,
// 				ProductType: "learning",
// 				Name: "สมุดเส้นเดียว 40 แผ่น",
// 				Price: 10,
// 				LearningSupplyID: 2,
// 			},
// 			{
// 				ProductCode: 8858784730205,
// 				ProductType: "learning",
// 				Name: "ดินน้ำมัน 150 กรัม",
// 				Price: 15,
// 				LearningSupplyID: 3,
// 			},
// 			{
// 				ProductCode: 8851907285142,
// 				ProductType: "learning",
// 				Name: "ดินสอดำ HB",
// 				Price: 6,
// 				LearningSupplyID: 4,
// 			},
// 			{
// 				ProductCode: 8851907171063,
// 				ProductType: "learning",
// 				Name: "ยางลบดินสอ",
// 				Price: 6,
// 				LearningSupplyID: 5,
// 			},
// 			{
// 				ProductCode: 8851907075057,
// 				ProductType: "learning",
// 				Name: "สีเทียน 12 สี มาสเตอร์อาร์ต",
// 				Price: 35,
// 				LearningSupplyID: 6,
// 			},
// 			{
// 				ProductCode: 8857121430597,
// 				ProductType: "learning",
// 				Name: "กระดาษ A4 (แพ็ค 40 แผ่น)",
// 				Price: 27,
// 				LearningSupplyID: 7,
// 			},
// 			{
// 				ProductCode: 8859694900088,
// 				ProductType: "book",
// 				Name: "คณิตคิดสนุกระดับอนุบาลเล่ม 1",
// 				Price: 90,
// 				BookID: 1,
// 			},
// 		}

// 		// Use repos with tx
// 		userRepoTx := NewGormUserRepo(tx)
// 		productRepoTx := NewGormProductRepo(tx)

// 		if err := userRepoTx.Insert(user); err != nil {
// 			return err
// 		}

// 		if err := productRepoTx.Insert(product); err != nil {
// 			return err
// 		}

// 		return nil
// 	})
// }

type InsertMultipleUseCase struct {
	MockRepo *repositories.MockRepository
	DB       *gorm.DB
}

func NewInsertMultipleUseCase(mockRepo *repositories.MockRepository, db *gorm.DB) *InsertMultipleUseCase {
	return &InsertMultipleUseCase{
		MockRepo: mockRepo,
		DB:       db,
	}
}

// func (uc *InsertMultipleUseCase) Execute() error {
// 	return uc.DB.Transaction(func(tx *gorm.DB) error {
// 		users := []domain.User{
// 			{Email: "nut1@mail.com", Password: utils.HashPassword("12341234"), Role: "ADMIN"},
// 			{Email: "nut2@mail.com", Password: utils.HashPassword("12341234"), Role: "USER"},
// 		}
// 		if err := tx.Create(&users).Error; err != nil {
// 			return err
// 		}

// 		books := []domain.Book{
// 			{Author: "พรพิไล", ISBN: "ISBN-001", Discount: 0},
// 			{Author: "พรพิไล", ISBN: "ISBN-002", Discount: 10},
// 			{Author: "พรพิไล", ISBN: "ISBN-003", Discount: 15},
// 			{Author: "พรพิไล", ISBN: "ISBN-004", Discount: 20},
// 		}
// 		if err := tx.Create(&books).Error; err != nil {
// 			return err
// 		}

// 		// สร้าง LearningSupplies
// 		learnings := []domain.LearningSupply{
// 			{Brand: "", Material: ""},
// 			{Brand: "", Material: ""},
// 			{Brand: "", Material: ""},
// 			{Brand: "", Material: ""},
// 			{Brand: "", Material: ""},
// 			{Brand: "", Material: ""},
// 			{Brand: "", Material: ""},
// 		}
// 		if err := tx.Create(&learnings).Error; err != nil {
// 			return err
// 		}
// 		products := []domain.Product{
// 			{
// 				ProductCode:      2021580100087,
// 				ProductType:      "learning",
// 				Name:             "ซองซิปตาข่าย A4",
// 				Price:            20,
// 				LearningSupplyID: &learnings[0].ID,
// 			},
// 			{
// 				ProductCode:      8851419014667,
// 				ProductType:      "learning",
// 				Name:             "สมุดเส้นเดียว 40 แผ่น",
// 				Price:            10,
// 				LearningSupplyID: &learnings[1].ID,
// 			},
// 			{
// 				ProductCode:      8858784730205,
// 				ProductType:      "learning",
// 				Name:             "ดินน้ำมัน 150 กรัม",
// 				Price:            15,
// 				LearningSupplyID: &learnings[2].ID,
// 			},
// 			{
// 				ProductCode:      8851907285142,
// 				ProductType:      "learning",
// 				Name:             "ดินสอดำ HB",
// 				Price:            6,
// 				LearningSupplyID: &learnings[3].ID,
// 			},
// 			{
// 				ProductCode:      8851907171063,
// 				ProductType:      "learning",
// 				Name:             "ยางลบดินสอ",
// 				Price:            6,
// 				LearningSupplyID: &learnings[4].ID,
// 			},
// 			{
// 				ProductCode:      8851907075057,
// 				ProductType:      "learning",
// 				Name:             "สีเทียน 12 สี มาสเตอร์อาร์ต",
// 				Price:            35,
// 				LearningSupplyID: &learnings[5].ID,
// 			},
// 			{
// 				ProductCode:      8857121430597,
// 				ProductType:      "learning",
// 				Name:             "กระดาษ A4 (แพ็ค 40 แผ่น)",
// 				Price:            27,
// 				LearningSupplyID: &learnings[6].ID,
// 			},
// 			{
// 				ProductCode: 8859694900088,
// 				ProductType: "book",
// 				Name:        "คณิตคิดสนุกระดับอนุบาลเล่ม 1",
// 				Price:       90,
// 				BookID:      &books[0].ID,
// 			},
// 			{
// 				ProductCode: 8859694900095,
// 				ProductType: "book",
// 				Name:        "คณิตคิดสนุกระดับอนุบาลเล่ม 2",
// 				Price:       90,
// 				BookID:      &books[1].ID,
// 			},
// 			{
// 				ProductCode: 8859694900101,
// 				ProductType: "book",
// 				Name:        "คณิตคิดสนุกระดับอนุบาลเล่ม 3",
// 				Price:       90,
// 				BookID:      &books[2].ID,
// 			},
// 			{
// 				ProductCode: 8899694900118,
// 				ProductType: "book",
// 				Name:        "คณิตคิดสนุกระดับอนุบาลเล่ม 4",
// 				Price:       90,
// 				BookID:      &books[3].ID,
// 			},
// 		}

// 		return uc.MockRepo.InsertMultiple(users, books, learnings, products)
// 	})
// }

func (uc *InsertMultipleUseCase) Execute() error {
	return uc.DB.Transaction(func(tx *gorm.DB) error {
		// 1. Insert Users
		users := []domain.User{
			{Email: "nut1@mail.com", Password: utils.HashPassword("12341234"), Role: "ADMIN"},
			{Email: "nut2@mail.com", Password: utils.HashPassword("12341234"), Role: "USER"},
		}
		if err := tx.Create(&users).Error; err != nil {
			return err
		}

		// 2. Insert Products ประเภท Learning ก่อน (ยังไม่เชื่อม LearningSupply)
		productsLearning := []domain.Product{
			{ProductCode: 2021580100087, ProductType: "learning", Name: "ซองซิปตาข่าย A4", Price: 20},
			{ProductCode: 8851419014667, ProductType: "learning", Name: "สมุดเส้นเดียว 40 แผ่น", Price: 10},
			{ProductCode: 8858784730205, ProductType: "learning", Name: "ดินน้ำมัน 150 กรัม", Price: 15},
			{ProductCode: 8851907285142, ProductType: "learning", Name: "ดินสอดำ HB", Price: 6},
			{ProductCode: 8851907171063, ProductType: "learning", Name: "ยางลบดินสอ", Price: 6},
			{ProductCode: 8851907075057, ProductType: "learning", Name: "สีเทียน 12 สี มาสเตอร์อาร์ต", Price: 35},
			{ProductCode: 8857121430597, ProductType: "learning", Name: "กระดาษ A4 (แพ็ค 40 แผ่น)", Price: 27},
		}
		if err := tx.Create(&productsLearning).Error; err != nil {
			return err
		}

		// 3. Insert LearningSupply พร้อม ProductID ที่ถูกต้อง
		learnings := []domain.LearningSupply{
			{Brand: "", Material: "", ProductID: productsLearning[0].ID},
			{Brand: "", Material: "", ProductID: productsLearning[1].ID},
			{Brand: "", Material: "", ProductID: productsLearning[2].ID},
			{Brand: "", Material: "", ProductID: productsLearning[3].ID},
			{Brand: "", Material: "", ProductID: productsLearning[4].ID},
			{Brand: "", Material: "", ProductID: productsLearning[5].ID},
			{Brand: "", Material: "", ProductID: productsLearning[6].ID},
		}
		if err := tx.Create(&learnings).Error; err != nil {
			return err
		}

		// 4. Insert Products ประเภท Book (ยังไม่เชื่อม Book)
		productsBooks := []domain.Product{
			{ProductCode: 8859694900088, ProductType: "book", Name: "คณิตคิดสนุกระดับอนุบาลเล่ม 1", Price: 90},
			{ProductCode: 8859694900095, ProductType: "book", Name: "คณิตคิดสนุกระดับอนุบาลเล่ม 2", Price: 90},
			{ProductCode: 8859694900101, ProductType: "book", Name: "คณิตคิดสนุกระดับอนุบาลเล่ม 3", Price: 90},
			{ProductCode: 8899694900118, ProductType: "book", Name: "คณิตคิดสนุกระดับอนุบาลเล่ม 4", Price: 90},
		}
		if err := tx.Create(&productsBooks).Error; err != nil {
			return err
		}

		// 5. Insert Books พร้อมเชื่อม ProductID
		books := []domain.Book{
			{Author: "พรพิไล", ISBN: "ISBN-001", Discount: 0, ProductID: productsBooks[0].ID},
			{Author: "พรพิไล", ISBN: "ISBN-002", Discount: 10, ProductID: productsBooks[1].ID},
			{Author: "พรพิไล", ISBN: "ISBN-003", Discount: 15, ProductID: productsBooks[2].ID},
			{Author: "พรพิไล", ISBN: "ISBN-004", Discount: 20, ProductID: productsBooks[3].ID},
		}
		if err := tx.Create(&books).Error; err != nil {
			return err
		}

		return nil
	})
}
