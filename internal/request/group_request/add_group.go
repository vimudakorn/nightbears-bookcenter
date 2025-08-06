package grouprequest

type AddNewGroupRequest struct {
	Name        string  `json:"name"`
	EduLevel    string  `json:"edu_level"`
	Description string  `json:"description"`
	SalePrice   float64 `json:"sale_price"`
}

type CreateGroupWithProductsRequest struct {
	Name        string            `json:"name"`
	EduLevel    string            `json:"edu_level"`
	Description string            `json:"description"`
	SalePrice   float64           `json:"sale_price"`
	Products    []ProductQuantity `json:"products"`
}

type ProductQuantity struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
