package tagrequest

type AddNewTagRequest struct {
	Name string `json:"name" validate:"required"`
}
