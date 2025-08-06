package tagrequest

type RenameTagRequest struct {
	Name string `json:"name" validate:"required,unique"`
}
