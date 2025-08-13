package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vimudakorn/internal/usecases"
)

type CartHandler struct {
	usecases *usecases.CartUsecase
}

func NewCartHandler(uc *usecases.CartUsecase) *CartHandler {
	return &CartHandler{usecases: uc}
}

func (h *CartHandler) GetByUserID(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user ID")
	}

	cart, err := h.usecases.GetCartByUserID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to find cart by user ID"})
	}
	return c.JSON(fiber.Map{
		"data": cart,
	})
}
