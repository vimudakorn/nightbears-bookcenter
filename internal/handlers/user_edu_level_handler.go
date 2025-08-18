package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vimudakorn/internal/usecases"
)

type UserEduLevelHandler struct {
	usecases *usecases.UserEduLevelUsecase
}

func NewUserEduLevelHandler(uc *usecases.UserEduLevelUsecase) *UserEduLevelHandler {
	return &UserEduLevelHandler{usecases: uc}
}

// POST /users/:id/edu-level
func (h *UserEduLevelHandler) CreateWithFixedLevels(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id type")
	}

	var req map[string]int // {"ป.1": 10, "ป.2": 5}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	edu, err := h.usecases.CreateWithFixedLevels(uint(userID), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(edu)
}

// PUT /users/:id/edu-level
func (h *UserEduLevelHandler) UpdateMultipleLevels(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id type")
	}

	var req map[string]int // {"ป.1": 12, "ป.3": 8}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := h.usecases.UpdateMultipleLevels(uint(userID), req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "updated successfully"})
}

// GET /users/:id/edu-level
func (h *UserEduLevelHandler) GetByUserID(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id type")
	}

	edu, err := h.usecases.GetByUserID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(edu)
}

// PUT /users/:id/edu-level
func (h *UserEduLevelHandler) UpdateStudentCount(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id type")
	}

	var body struct {
		LevelName string `json:"level_name"`
		Count     int    `json:"count"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
	}

	if err := h.usecases.UpdateStudentCount(uint(userID), body.LevelName, body.Count); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "updated successfully"})
}

// DELETE /users/:id/edu-level
func (h *UserEduLevelHandler) DeleteByUserID(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id type")
	}

	if err := h.usecases.DeleteByUserID(uint(userID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "deleted successfully"})
}

// func (h *UserEduLevelHandler) AddEduLevel(c *fiber.Ctx) error {
// 	userID, ok := c.Locals("user_id").(uint)
// 	if !ok {
// 		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id type")
// 	}

// 	var req useredulevelrequest.AddMultiEduLevelRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
// 	}

// 	var eduLevelsToCreate []domain.UserEduLevel
// 	for _, edu := range req.Levels {
// 		if edu.StudentCount < 1 {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "StudentCount must be at least 1"})
// 		}

// 		isExist, err := h.usecases.IsEduLevelNameExist(edu.EduLevel, edu.EduYear, userID)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to check education level"})
// 		}
// 		if isExist {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("EduLevel %s already exists", edu.EduLevel)})
// 		}

// 		eduLevelsToCreate = append(eduLevelsToCreate, domain.UserEduLevel{
// 			UserID:       userID,
// 			EduLevel:     edu.EduLevel,
// 			StudentCount: edu.StudentCount,
// 			EduYear:      edu.EduYear,
// 		})
// 	}

// 	if err := h.usecases.CreateMultiple(eduLevelsToCreate); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create education levels"})
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Education levels added successfully"})
// }

// func (h *UserEduLevelHandler) UpdateEduLevels(c *fiber.Ctx) error {
// 	userID, ok := c.Locals("user_id").(uint)
// 	if !ok {
// 		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id type")
// 	}

// 	var req useredulevelrequest.UpdateMultiEduLevelRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
// 	}

// 	var levels []domain.UserEduLevel
// 	for _, l := range req.Levels {
// 		if l.StudentCount < 1 {
// 			return fiber.NewError(fiber.StatusBadRequest, "student_count must be at least 1")
// 		}
// 		levels = append(levels, domain.UserEduLevel{
// 			UserID:       userID,
// 			EduLevel:     l.EduLevel,
// 			EduYear:      l.EduYear,
// 			StudentCount: l.StudentCount,
// 		})
// 	}

// 	if err := h.usecases.UpdateMultiple(userID, levels); err != nil {
// 		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
// 	}

// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "education levels updated"})
// }

// func (h *UserEduLevelHandler) UpdateEduLevel(c *fiber.Ctx) error {
// 	eduID, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user education level ID"})
// 	}
// 	var req useredulevelrequest.UpdateEduLevelRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, "Invalid request")
// 	}

// 	updated := &domain.UserEduLevel{
// 		EduLevel:     req.EduLevel,
// 		StudentCount: req.StudentCount,
// 		EduYear:      req.EduYear,
// 	}

// 	if err := h.usecases.UpdateByID(uint(eduID), updated); err != nil {
// 		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update edu level")
// 	}

// 	return c.SendStatus(fiber.StatusOK)
// }

// func (h *UserEduLevelHandler) Delete(c *fiber.Ctx) error {
// 	eduLevelID, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user education level ID"})
// 	}
// 	isEduLevelIDExist, err := h.usecases.IsEduLevelExist(uint(eduLevelID))
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to check edu level ID"})
// 	}
// 	if !isEduLevelIDExist {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User Education Level ID not found"})
// 	}

// 	if err := h.usecases.Delete(uint(eduLevelID)); err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete edu level from this user"})
// 	}
// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{
// 		"message": "Education Level deleted from user successfully",
// 	})
// }

// func (h *UserEduLevelHandler) GetByUserID(c *fiber.Ctx) error {
// 	userID, ok := c.Locals("user_id").(uint)
// 	if !ok {
// 		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id type")
// 	}

// 	levels, err := h.usecases.GetByUserID(userID)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get education level by id"})
// 	}
// 	return c.JSON(fiber.Map{
// 		"data": levels,
// 	})

// }
