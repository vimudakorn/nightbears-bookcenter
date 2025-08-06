package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vimudakorn/internal/domain"
	tagrequest "github.com/vimudakorn/internal/request/tag_request"
	"github.com/vimudakorn/internal/usecases"
	"github.com/vimudakorn/internal/utils"
)

type TagHandler struct {
	usecases *usecases.TagUsecase
}

func NewTagHandler(uc *usecases.TagUsecase) *TagHandler {
	return &TagHandler{usecases: uc}
}

func (h *TagHandler) GetAll(c *fiber.Ctx) error {
	name := c.Query("name")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortBy := c.Query("sortBy", "name")
	sortOrder := c.Query("sortOrder", "asc")
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}
	tags, count, err := h.usecases.GetPagination(page, limit, name, sortBy, sortOrder)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{
		"full": tags,
		// "data":       res,
		"page":       page,
		"limit":      limit,
		"count":      count,
		"totalPages": (int(count) + limit - 1) / limit,
	})
}

func (h *TagHandler) AddNewTag(c *fiber.Ctx) error {
	var req tagrequest.AddNewTagRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := utils.ValidateStruct(c, &req); err != nil {
		return err
	}

	isExist, err := h.usecases.IsNameExists(req.Name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal error checking name"})
	}
	if isExist {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Tag name already exists"})
	}

	tag := &domain.Tag{
		Name: req.Name,
	}

	if err := h.usecases.Create(tag); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create tag"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Tag created successfully"})
}

func (h *TagHandler) RenameTag(c *fiber.Ctx) error {
	tagID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid tag id"})
	}

	var req tagrequest.RenameTagRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := utils.ValidateStruct(c, &req); err != nil {
		return err
	}

	if err := h.usecases.RenameTag(uint(tagID), req.Name); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to rename tag",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Tag renamed successfully",
	})
}

func (h *TagHandler) Delete(c *fiber.Ctx) error {
	tagID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid tag id"})
	}

	hasUsed, err := h.usecases.IsTagHasUsed(uint(tagID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to check related tags"})
	}

	if hasUsed {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error":   "this tag is being used by some products",
			"hasUsed": true,
			"message": "Cannot delete tag that is being used by products.",
		})
	}

	if err = h.usecases.Delete(uint(tagID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete tag info"})
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "Tag deleted successfully",
	})
}
