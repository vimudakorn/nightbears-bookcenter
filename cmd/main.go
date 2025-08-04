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
	authUsecase := usecases.NewAuthUsecase(authRepo)
	authHandler := handlers.NewAuthHandler(authUsecase)

	userRepo := repositories.NewUserGormRepo(db)
	userUsecase := usecases.NewUserUsecase(userRepo)
	userHandler := handlers.NewUserHandler(userUsecase)

	productRepo := repositories.NewProductGormDB(db)
	productUsecase := usecases.NewProductUsecase(productRepo)
	productHandler := handlers.NewProductHandler(productUsecase)

	app := fiber.New()
	routes.SetupRoutes(app, authHandler, userHandler, productHandler)
	app.Listen(":8080")
}
