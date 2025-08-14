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
	groupHandler *handlers.GroupHandler,
	groupProductHandler *handlers.GroupProductHandler,
	cartHandler *handlers.CartHandler,
	cartItemHandler *handlers.CartItemHandler,
	orderHandler *handlers.OrderHandler,
	orderItemHandler *handlers.OrderItemHandler,
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
	// Put: All Fields
	// Patch: Some Fields

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

	api.Get("/tags", tagHandler.GetAll)
	api.Post("/tags", tagHandler.AddNewTag)
	api.Put("/tags/:id/rename", tagHandler.RenameTag)
	api.Delete("/tags/:id", tagHandler.Delete)

	api.Get("/groups", groupHandler.GetAll)
	api.Post("/groups/add-product", groupHandler.CreateGroupWithProducts)
	api.Post("/groups", groupHandler.AddNewGroup)
	api.Put("/groups/:id", groupHandler.Update)
	api.Delete("/groups/:id", groupHandler.Delete)

	api.Get("/groups/:id", groupProductHandler.GetByID)
	api.Post("/groups/:id/add-product", groupProductHandler.AddMultiProductInGroup)
	api.Post("/groups/:id/add-product-tx", groupProductHandler.AddMultiProductInGroupWithTx)
	api.Patch("/groups/:id", groupProductHandler.UpdateProductInGroupID)
	api.Patch("/groups/:id/products/:productID", groupProductHandler.UpdateProductInGroup)
	api.Delete("/groups/:id/products/:productID", groupProductHandler.DeleteProductInGroup)

	api.Get("/carts/:user_id", cartHandler.GetByUserID)

	// there are func
	// 1. AddOrUpdateCartItem (1)
	// 2. UpdateItemInCart (1)
	api.Get("/carts", cartItemHandler.GetOwnItemInCard)
	api.Post("/carts", cartItemHandler.AddOrUpdateMultiCartItems)
	api.Patch("/carts/:id", cartItemHandler.Update)
	api.Delete("/carts/:cartItemID", cartItemHandler.DeleteCartItem)

	api.Post("/orders", orderHandler.Create)
	api.Get("/orders", orderHandler.GetAll)
	api.Get("/orders/:id", orderHandler.GetByID)
	api.Put("/orders/:id", orderHandler.Update)
	api.Delete("/orders/:id", orderHandler.Delete)
}
