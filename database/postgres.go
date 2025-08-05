package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/vimudakorn/configs"
	"github.com/vimudakorn/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	port, err := strconv.Atoi(configs.GetEnv("DB_PORT", "5432"))
	if err != nil {
		log.Fatal("Invalid DB_PORT")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", configs.GetEnv("DB_HOST", "localhost"), port, configs.GetEnv("DB_USER", "postgres"), configs.GetEnv("DB_PASSWORD", ""), configs.GetEnv("DB_NAME", "postgres"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	DB = db
	return DB
}

func Migrate(db *gorm.DB) {
	// db.Migrator().DropTable(&domain.User{}, &domain.Profile{}, &domain.Cart{}, &domain.Category{}, &domain.Product{}, &domain.Book{}, &domain.LearningSupply{}, &domain.OfficeSupply{}, &domain.Order{}, &domain.OrderItem{}, &domain.CartItem{}, &domain.ProductCategory{}, &domain.ProductImage{}, &domain.Group{}, &domain.GroupProduct{})

	// if err := db.AutoMigrate(&models.User{}, &models.Profile{}, &models.Cart{}, &models.Category{}, &models.Product{}, &models.Book{}, &models.LearningSupply{}, &models.OfficeSupply{}, &models.Order{}, &models.OrderItem{}, &models.CartItem{}, &models.BookCategory{}, &models.BookImage{}, &models.Group{}, &models.GroupProduct{}); err != nil {
	if err := db.AutoMigrate(&domain.User{}, &domain.Profile{}, &domain.Cart{}, &domain.Category{}, &domain.Product{}, &domain.Book{}, &domain.LearningSupply{}, &domain.OfficeSupply{}, &domain.Order{}, &domain.OrderItem{}, &domain.CartItem{}, &domain.ProductCategory{}, &domain.ProductImage{}, &domain.Group{}, &domain.GroupProduct{}, &domain.Tag{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// mockRepo := repositories.NewMockRepository(db)
	// insertUseCase := usecases.NewInsertMultipleUseCase(mockRepo, db)

	// if err := insertUseCase.Execute(); err != nil {
	// 	log.Fatalf("Insert data failed: %v", err)
	// } else {
	// 	log.Println("Insert data success!")
	// }

}
