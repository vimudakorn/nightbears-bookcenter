package groupproductrequest

type AddProductInGroupRequest struct {
	// GroupID   uint `json:"group_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
