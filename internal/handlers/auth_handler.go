package handlers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vimudakorn/constants"
	"github.com/vimudakorn/internal/request"
	"github.com/vimudakorn/internal/usecases"
	"github.com/vimudakorn/internal/utils"
)

type AuthHandler struct {
	usecases *usecases.AuthUsecase
}

func NewAuthHandler(uc *usecases.AuthUsecase) *AuthHandler {
	return &AuthHandler{usecases: uc}
}

// func (h *AuthHandler) Register(c *fiber.Ctx) error {
// 	var req request.RegisterRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
// 	}

// 	if err := utils.ValidateStruct(c, &req); err != nil {
// 		return err
// 	}

// 	if err := h.usecases.Register(&req); err != nil {
// 		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
// 	}
// 	return c.JSON(fiber.Map{"message": "registered"})

// }

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req request.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	warnings, err := h.usecases.Register(&req)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if len(warnings) > 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":  "User created with warnings",
			"warnings": warnings,
		})
	}

	return c.JSON(fiber.Map{
		"message": "User created successfully",
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var body request.LoginRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
	}

	user, err := h.usecases.Login(body.Email, body.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "login failed"})
	}

	// Create tokens
	accessToken, _ := utils.GenerateToken(user.ID, user.Role, utils.AccessSecret, 15*time.Minute)
	refreshToken, _ := utils.GenerateToken(user.ID, user.Role, utils.RefreshSecret, 7*24*time.Hour)

	// Set refresh token in cookie
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/",
		HTTPOnly: true,
		// Secure:   os.Getenv("APP_ENV") == "production",
		SameSite: "Lax",
		MaxAge:   7 * 24 * 60 * 60,
	})

	return c.JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken, // optional: remove in production
	})
}

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return c.Status(401).JSON(fiber.Map{"error": "missing refresh token"})
	}

	claims, err := utils.ParseToken(refreshToken, utils.RefreshSecret)
	if err != nil {
		return c.Status(403).JSON(fiber.Map{"error": "invalid refresh token"})
	}

	accessToken, err := utils.GenerateToken(claims.UserID, claims.Role, utils.AccessSecret, 15*time.Minute)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "cannot generate access token"})
	}

	return c.JSON(fiber.Map{"access_token": accessToken})
}

func (h *AuthHandler) ChangePassword(c *fiber.Ctx) error {
	var req request.ChangePasswordRequest
	userID := c.Locals("user_id").(uint)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := utils.ValidateStruct(c, &req); err != nil {
		return err
	}

	if err := h.usecases.ChangePassword(userID, req.OldPassword, req.NewPassword, req.ConfirmNewPassword); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Password changed successfully",
	})
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	idParam, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user id"})
	}

	requestingUserID := c.Locals("user_id").(uint)
	requestingUserRole := c.Locals("role").(string)

	if requestingUserRole != constants.ADMIN && requestingUserID != uint(idParam) {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "you are not authorized to delete this user",
		})
	}

	if err := h.usecases.Delete(uint(idParam)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *AuthHandler) Me(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(float64)
	role := c.Locals("role").(string)

	return c.JSON(fiber.Map{
		"user_id": userID,
		"role":    role,
	})
}
