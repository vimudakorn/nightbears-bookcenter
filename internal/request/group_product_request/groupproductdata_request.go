package groupproductrequest

type GroupProductWithDetail struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    uint    `json:"quantity"`
	GroupID     uint    `json:"group_id"`
	GroupName   string  `json:"group_name"`
}
