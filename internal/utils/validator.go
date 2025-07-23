package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func ValidateStruct(c *fiber.Ctx, s interface{}) error {
	if err := validate.Struct(s); err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			errorMap := make(map[string]string)
			for _, e := range validationErrs {
				field := e.Field()
				tag := e.Tag()
				errorMap[field] = fmt.Sprintf("validation failed on '%s'", tag)
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"validation_errors": errorMap,
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}
	return nil
}
