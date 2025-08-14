package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vimudakorn/internal/domain"
	useredulevelrequest "github.com/vimudakorn/internal/request/user_edu_level_request"
	"github.com/vimudakorn/internal/usecases"
)

type UserEduLevelHandler struct {
	usecases *usecases.UserEduLevelUsecase
}

func NewUserEduLevelHandler(uc *usecases.UserEduLevelUsecase) *UserEduLevelHandler {
	return &UserEduLevelHandler{usecases: uc}
}

func (h *UserEduLevelHandler) AddEduLevel(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id type")
	}

	var req useredulevelrequest.AddMultiEduLevelRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	var eduLevelsToCreate []domain.UserEduLevel
	for _, edu := range req.Levels {
		if edu.StudentCount < 1 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "StudentCount must be at least 1"})
		}

		isExist, err := h.usecases.IsEduLevelNameExist(edu.EduLevel, edu.EduYear, userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to check education level"})
		}
		if isExist {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("EduLevel %s already exists", edu.EduLevel)})
		}

		eduLevelsToCreate = append(eduLevelsToCreate, domain.UserEduLevel{
			UserID:       userID,
			EduLevel:     edu.EduLevel,
			StudentCount: edu.StudentCount,
			EduYear:      edu.EduYear,
		})
	}

	if err := h.usecases.CreateMultiple(eduLevelsToCreate); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create education levels"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Education levels added successfully"})
}

func (h *UserEduLevelHandler) UpdateEduLevels(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id type")
	}

	var req useredulevelrequest.UpdateMultiEduLevelRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	var levels []domain.UserEduLevel
	for _, l := range req.Levels {
		if l.StudentCount < 1 {
			return fiber.NewError(fiber.StatusBadRequest, "student_count must be at least 1")
		}
		levels = append(levels, domain.UserEduLevel{
			UserID:       userID,
			EduLevel:     l.EduLevel,
			EduYear:      l.EduYear,
			StudentCount: l.StudentCount,
		})
	}

	if err := h.usecases.UpdateMultiple(userID, levels); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "education levels updated"})
}

func (h *UserEduLevelHandler) UpdateEduLevel(c *fiber.Ctx) error {
	var req useredulevelrequest.UpdateEduLevelRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	}

	updated := &domain.UserEduLevel{
		EduLevel:     req.EduLevel,
		StudentCount: req.StudentCount,
		EduYear:      req.EduYear,
	}

	if err := h.usecases.UpdateByID(req.ID, updated); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update edu level")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *UserEduLevelHandler) Delete(c *fiber.Ctx) error {
	eduLevelID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user education level ID"})
	}
	isEduLevelIDExist, err := h.usecases.IsEduLevelExist(uint(eduLevelID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to check edu level ID"})
	}
	if !isEduLevelIDExist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User Education Level ID not found"})
	}

	if err := h.usecases.Delete(uint(eduLevelID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete edu level from this user"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Education Level deleted from user successfully",
	})
}

func (h *UserEduLevelHandler) GetByUserID(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id type")
	}

	levels, err := h.usecases.GetByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get education level by id"})
	}
	return c.JSON(fiber.Map{
		"data": levels,
	})

}
