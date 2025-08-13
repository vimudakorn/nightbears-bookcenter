package cartitemresponse

type CartItemDetailResponse struct {
	CartID      uint    `json:"cart_id"`
	UserID      uint    `json:"user_id"`
	UserEmail   string  `json:"user_email"`
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
	Quantity    uint    `json:"quantity"`
}
