package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vimudakorn/internal/handlers"
	"github.com/vimudakorn/internal/middleware"
)

func SetupRoutes(app *fiber.App,
	authHandler *handlers.AuthHandler,
	userHandler *handlers.UserHandler,
	productHandler *handlers.ProductHandler,
	categoryHandler *handlers.CategoryHandler,
	tagHandler *handlers.TagHandler,
) {
	// Public routes
	app.Post("/login", authHandler.Login)
	app.Post("/register", authHandler.Register)
	app.Post("/refresh-token", authHandler.RefreshToken)

	api := app.Group("/api", middleware.JWTMiddleware)

	api.Get("/me", func(c *fiber.Ctx) error {
		userID := c.Locals("user_id")
		role := c.Locals("role")

		return c.JSON(fiber.Map{
			"message": "Welcome to api/me",
			"user_id": userID,
			"role":    role,
		})
	})

	api.Get("/users", userHandler.GetAll)
	api.Put("/users/change-profile", userHandler.ChangeProfileData)
	api.Put("/users/change-password", authHandler.ChangePassword)

	api.Get("/products", productHandler.GetAll)
	api.Post("/products", productHandler.AddNewProduct)
	api.Put("/products/:id", productHandler.UpdateProduct)
	api.Delete("/products/:id", productHandler.Delete)

	api.Get("/category", categoryHandler.GetAll)
	api.Post("/category", categoryHandler.CreateCategory)
	api.Put("/category/:id", categoryHandler.Update)
	api.Delete("/category/:id", categoryHandler.Delete)
}
