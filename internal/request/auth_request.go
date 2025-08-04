package request

type RegisterRequest struct {
	Name            string `json:"name"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
	Role            string `json:"role" validate:"required,oneof=admin user"`
	Phone           string `json:"phone" validate:"required,len=10"`
	Address         string `json:"address"`
	AvatarURL       string `json:"avatar_url"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
