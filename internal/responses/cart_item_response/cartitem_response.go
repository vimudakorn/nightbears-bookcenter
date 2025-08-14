package cartitemresponse

// type CartItemDetailResponse struct {
// 	CartID      uint    `json:"cart_id"`
// 	UserID      uint    `json:"user_id"`
// 	UserEmail   string  `json:"user_email"`
// 	ProductID   uint    `json:"product_id"`
// 	ProductName string  `json:"product_name"`
// 	Price       float64 `json:"price"`
// 	Quantity    uint    `json:"quantity"`
// }

type CartItemDetailResponse struct {
	CartID       uint     `json:"cart_id"`
	UserID       uint     `json:"user_id"`
	UserEmail    string   `json:"user_email"`
	ProductID    *uint    `json:"product_id,omitempty"`
	ProductName  *string  `json:"product_name,omitempty"`
	ProductPrice *float64 `json:"product_price,omitempty"`
	GroupID      *uint    `json:"group_id,omitempty"`
	GroupName    *string  `json:"group_name,omitempty"`
	GroupPrice   *float64 `json:"group_price,omitempty"`
	Quantity     int      `json:"quantity"`
}
