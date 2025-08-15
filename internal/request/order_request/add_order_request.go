package orderrequest

type CreateOrderRequest struct {
	UserID     uint                 `json:"user_id"`
	TotalPrice float64              `json:"total_price"`
	Status     string               `json:"status"`
	Items      []CreateOrderItemReq `json:"items"`
}

type CreateOrderItemReq struct {
	ProductID *uint   `json:"product_id,omitempty"`
	GroupID   *uint   `json:"group_id,omitempty"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
