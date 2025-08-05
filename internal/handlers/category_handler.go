package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vimudakorn/internal/domain"
	categoryrequest "github.com/vimudakorn/internal/request/category_request"
	"github.com/vimudakorn/internal/usecases"
)

type CategoryHandler struct {
	usecases *usecases.CategoryUsecase
}

func NewCategoryHandler(uc *usecases.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{usecases: uc}
}

// func (h *CategoryHandler) GetAll(c *fiber.Ctx) error {
// 	var
// }

func (h *CategoryHandler) CreateCategory(c *fiber.Ctx) error {
	var req categoryrequest.AddNewCategoryRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// ตรวจสอบว่าชื่อนี้มีอยู่แล้วหรือยัง
	isExist, err := h.usecases.IsNameExists(req.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal error checking name"})
	}
	if isExist {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Category name already exists"})
	}

	isParentExist, err := h.usecases.IsParentExist(req.ParentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal error checking name"})
	}
	if !isParentExist {
		return c.Status(400).JSON(fiber.Map{"error": "Parent category does not exist"})
	}

	category := &domain.Category{
		Name:     req.Name,
		ParentID: req.ParentID,
	}

	if err := h.usecases.Create(category); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create category")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Category created successfully"})
}
