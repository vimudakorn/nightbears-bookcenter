package categoryrequest

type UpdateCategoryRequest struct {
	Name     string `json:"name"`
	ParentID *uint  `json:"parent_id"`
}
