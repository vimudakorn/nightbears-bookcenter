package grouprequest

type AddNewGroupRequest struct {
	Name        string  `json:"name"`
	EduLevel    string  `json:"edu_level"`
	Description string  `json:"description"`
	SalePrice   float64 `json:"sale_price"`
}
