package main

import (
	"github.com/vimudakorn/configs"
	"github.com/vimudakorn/database"
	"github.com/vimudakorn/internal/handlers"
	"github.com/vimudakorn/internal/repositories"
	"github.com/vimudakorn/internal/routes"
	"github.com/vimudakorn/internal/usecases"

	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.LoadEnv()
	db := database.ConnectDB()
	database.Migrate(db)

	authRepo := repositories.NewUserGormRepo(db)
	userRepo := repositories.NewUserGormRepo(db)
	tagRepo := repositories.NewTagGormRepo(db)
	productRepo := repositories.NewProductGormDB(db)
	categoryRepo := repositories.NewCategoryGormRepo(db)
	groupProductRepo := repositories.NewGroupProductGormRepo(db)
	groupRepo := repositories.NewGroupGormRepo(db)

	authUsecase := usecases.NewAuthUsecase(authRepo)
	userUsecase := usecases.NewUserUsecase(userRepo)
	tagUsecase := usecases.NewTagUsecase(tagRepo)
	productUsecase := usecases.NewProductUsecase(productRepo, tagRepo)
	categoryUsecase := usecases.NewCategoryUsecase(categoryRepo)
	groupProductUsecase := usecases.NewGroupProductUsecase(groupProductRepo, groupRepo)
	groupUsecase := usecases.NewGroupUsecase(groupRepo, groupProductRepo)

	authHandler := handlers.NewAuthHandler(authUsecase)
	userHandler := handlers.NewUserHandler(userUsecase)
	tagHandler := handlers.NewTagHandler(tagUsecase)
	productHandler := handlers.NewProductHandler(productUsecase)
	categoryHandler := handlers.NewCategoryHandler(categoryUsecase)
	groupProductHandler := handlers.NewGroupProductUsecase(groupProductUsecase)
	groupHandler := handlers.NewGroupHandler(groupUsecase)

	// mockUsecase := usecases.NewMockUseCase(db, userRepo, productRepo)
	// mockUsecase.AddNewData()

	app := fiber.New()
	routes.SetupRoutes(app, authHandler, userHandler, productHandler, categoryHandler, tagHandler, groupHandler, groupProductHandler)
	app.Listen(":8080")
}
