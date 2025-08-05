package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vimudakorn/constants"
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

func (h *CategoryHandler) GetAll(c *fiber.Ctx) error {
	categories, err := h.usecases.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err})
	}
	return c.JSON(fiber.Map{
		"data": categories})
}

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

func (h *CategoryHandler) Update(c *fiber.Ctx) error {
	categoryID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid category id"})
	}

	var req categoryrequest.UpdateCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	category, err := h.usecases.FindByID(uint(categoryID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "category not found"})
	}

	if req.Name != category.Name && req.Name != "" {
		category.Name = req.Name
	}
	if req.ParentID != nil {
		if *req.ParentID == category.ID {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "category cannot be its own parent",
			})
		}

		exists, err := h.usecases.IsIDExists(req.ParentID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to validate parent_id",
			})
		}
		if !exists {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "parent category does not exist",
			})
		}
		category.ParentID = req.ParentID
	} else {
		category.ParentID = nil
	}

	if err := h.usecases.Update(category); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update category"})
	}

	return c.JSON(fiber.Map{
		"message":  "category updated successfully",
		"category": category,
	})
}

func (h *CategoryHandler) Delete(c *fiber.Ctx) error {
	categoryID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid category id"})
	}

	requestingUserRole := c.Locals("role").(string)
	if requestingUserRole != constants.ADMIN {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "you are not authorized to delete this product",
		})
	}

	_, err = h.usecases.FindByID(uint(categoryID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "category not found"})
	}

	// if there are products using this category
	hasProducts, err := h.usecases.HasProducts(uint(categoryID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to check related products"})
	}

	if hasProducts {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error":       "this category is being used by some products",
			"hasProducts": true,
			"message":     "Cannot delete category that is being used by products.",
		})
	}

	hasChildren, err := h.usecases.HasChildren(uint(categoryID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to check children"})
	}

	// ถ้ามีลูก และไม่มี query param "force=true" ให้แจ้งเตือน
	forceDelete := c.Query("force") == "true"
	if hasChildren && !forceDelete {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error":       "this category has subcategories",
			"hasChildren": true,
			"message":     "Are you sure you want to delete? This category has subcategories.",
		})
	}

	if err = h.usecases.Delete(uint(categoryID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete category info"})
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "category deleted successfully",
	})
}
