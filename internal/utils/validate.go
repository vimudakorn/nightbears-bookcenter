package utils

import (
	"strings"

	"github.com/vimudakorn/internal/request"
)

type Warning struct {
	Field   string
	Message string
}

func ValidateRegisterForm(req *request.RegisterRequest) []Warning {
	var warnings []Warning

	if req.Email == "" {
		warnings = append(warnings, Warning{Field: "email", Message: "Email is required"})
	}
	if req.Password == "" || req.ConfirmPassword == "" {
		warnings = append(warnings, Warning{Field: "password", Message: "Password and confirm password are required"})
	} else if req.Password != req.ConfirmPassword {
		warnings = append(warnings, Warning{Field: "password", Message: "Passwords do not match"})
	}
	if strings.TrimSpace(req.Name) == "" {
		warnings = append(warnings, Warning{Field: "name", Message: "Name is recommended"})
	}
	return warnings
}
