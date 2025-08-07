package groupproductrequest

type UpdateProductInGroupIDRequest struct {
	Products []ProductInGroup `json:"products"`
}

type ProductInGroup struct {
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
