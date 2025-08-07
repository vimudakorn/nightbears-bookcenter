package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vimudakorn/internal/domain"
	groupproductrequest "github.com/vimudakorn/internal/request/group_product_request"
	"github.com/vimudakorn/internal/usecases"
)

type GroupProductHandler struct {
	usecases *usecases.GroupProductUsecase
}

func NewGroupProductUsecase(uc *usecases.GroupProductUsecase) *GroupProductHandler {
	return &GroupProductHandler{usecases: uc}
}

func (h *GroupProductHandler) AddProductInGroup(c *fiber.Ctx) error {
	groupID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	isIDExist, err := h.usecases.IsGroupIDExists(uint(groupID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to check group id"})
	}
	if !isIDExist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Group ID not found"})
	}
	var req groupproductrequest.AddProductInGroupRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if req.Quantity < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Quantity must be at least 1",
		})
	}
	groupProduct := &domain.GroupProduct{
		GroupID:   uint(groupID),
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	group, err := h.usecases.AddProductInGroup(groupProduct)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to add product to group",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    group,
		"message": "Product added to group successfully",
	})
}
