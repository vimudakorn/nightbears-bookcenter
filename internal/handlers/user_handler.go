package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vimudakorn/constants"
	"github.com/vimudakorn/internal/request"
	"github.com/vimudakorn/internal/responses"
	"github.com/vimudakorn/internal/usecases"
	"github.com/vimudakorn/internal/utils"
)

type UserHandler struct {
	usecases *usecases.UserUsecase
}

func NewUserHandler(uc *usecases.UserUsecase) *UserHandler {
	return &UserHandler{usecases: uc}
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	userRole := c.Locals("role")
	name := c.Query("username")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortBy := c.Query("sortBy", "name")
	sortOrder := c.Query("sortOrder", "asc")

	if page < 1 {
		page = 1
	}

	users, count, err := h.usecases.GetPagination(page, limit, name, sortBy, sortOrder)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	switch userRole {
	case constants.ADMIN:
		var res []responses.AdmidUserdataResponse
		for _, user := range users {
			res = append(res, responses.AdmidUserdataResponse{
				ID:       user.ID,
				Name:     user.Name,
				Email:    user.Email,
				Password: user.Password,
				Role:     user.Role,
				Cart:     &user.Cart,
				Orders:   user.Orders,
			})
		}
		return c.JSON(fiber.Map{
			"data":       res,
			"page":       page,
			"limit":      limit,
			"count":      count,
			"totalPages": (int(count) + limit - 1) / limit,
		})
	case constants.USER:
		var res []responses.UserUserdataResponse
		for _, user := range users {
			res = append(res, responses.UserUserdataResponse{
				ID:   user.ID,
				Name: user.Name})
		}
		return c.JSON(fiber.Map{
			"data":       res,
			"page":       page,
			"limit":      limit,
			"count":      count,
			"totalPages": (int(count) + limit - 1) / limit,
		})
	default:
		return c.Status(403).JSON(fiber.Map{"error": "forbidden"})
	}
}

func (h *UserHandler) ChangeProfileData(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var req request.UserProfileDataRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := utils.ValidateStruct(c, &req); err != nil {
		return err
	}

	user, err := h.usecases.GetByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	hasChange := false
	if req.Name != "" && req.Name != user.Name {
		user.Name = req.Name
		hasChange = true
	}
	if req.Phone != "" && req.Phone != user.Phone {
		user.Phone = req.Phone
		hasChange = true
	}

	if !hasChange {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No changes detected",
		})
	}

	// ทำการ update
	if err := h.usecases.UpdateUserProfile(userID, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update user",
		})
	}

	return c.JSON(fiber.Map{"message": "Profile updated successfully"})

}
