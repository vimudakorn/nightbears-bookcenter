package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vimudakorn/internal/domain"
	errorsres "github.com/vimudakorn/internal/domain/errors_res"
	grouprequest "github.com/vimudakorn/internal/request/group_request"
	"github.com/vimudakorn/internal/usecases"
)

type GroupHandler struct {
	usecases *usecases.GroupUsecase
}

func NewGroupHandler(uc *usecases.GroupUsecase) *GroupHandler {
	return &GroupHandler{usecases: uc}
}

func (h *GroupHandler) GetPagination(c *fiber.Ctx) error {
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
	groups, count, err := h.usecases.GetPagination(page, limit, name, sortBy, sortOrder)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{
		"full": groups,
		// "data":       res,
		"page":       page,
		"limit":      limit,
		"count":      count,
		"totalPages": (int(count) + limit - 1) / limit,
	})
}

func (h *GroupHandler) AddNewGroup(c *fiber.Ctx) error {
	var req grouprequest.AddNewGroupRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	// name & edu_level can not be same display: <name> <edu_level>
	isExist, err := h.usecases.IsNameAndEduExist(req.Name, req.EduLevel)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal error checking name and education level"})
	}
	if isExist {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Category name already exists"})
	}

	group := &domain.Group{
		Name:        req.Name,
		EduLevel:    req.EduLevel,
		Description: req.Description,
		SalePrice:   req.SalePrice,
	}

	if err := h.usecases.Create(group); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create group")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Group created successfully"})
}

func (h *GroupHandler) Update(c *fiber.Ctx) error {
	groupID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}

	var req grouprequest.UpdateGroupRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err = h.usecases.UpdateGroup(uint(groupID), req)
	if err != nil {
		switch err {
		case errorsres.ErrGroupNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Group not found"})
		case errorsres.ErrGroupNameExist:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Group name with this education level already exists"})
		case errorsres.ErrInvalidSalePrice:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid sale price"})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal server error"})
		}
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *GroupHandler) Delete(c *fiber.Ctx) error {
	groupID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid group ID",
		})
	}

	// ถ้ามีคนเอา group นี้ไปใช้ใน cart , order เราต้องจัดการยังไง ลบในตะกร้าไปด้วยใช่มั้ย
	if err := h.usecases.Delete(uint(groupID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete group info"})
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"message": "group deleted successfully",
	})
}
