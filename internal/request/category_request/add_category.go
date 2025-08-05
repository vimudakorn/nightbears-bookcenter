package categoryrequest

type AddNewCategoryRequest struct {
	Name     string `json:"name"`
	ParentID *uint  `json:"parent_id"`
}
