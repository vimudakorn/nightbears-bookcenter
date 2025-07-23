package request

type UserProfileDataRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone" validate:"len=10"`
}

type ChangePasswordRequest struct {
	OldPassword        string `json:"old_password" validate:"required,min=8"`
	NewPassword        string `json:"new_password" validate:"required,min=8"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required"`
}
