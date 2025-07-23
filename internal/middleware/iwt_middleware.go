package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/vimudakorn/internal/utils"
)

func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(401).JSON(fiber.Map{"error": "missing or invalid token"})
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := utils.ParseToken(tokenStr, utils.AccessSecret)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "invalid or expired token"})
	}

	c.Locals("user_id", claims.UserID)
	c.Locals("role", claims.Role)
	return c.Next()
}

func RequireRoles(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
		if !ok {
			return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
		}

		for _, allowed := range allowedRoles {
			if strings.EqualFold(role, allowed) {
				return c.Next()
			}
		}

		return c.Status(403).JSON(fiber.Map{"error": "forbidden"})
	}
}
