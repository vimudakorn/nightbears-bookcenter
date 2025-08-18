package orderrequest

type UpdateOrderRequest struct {
	Status string                   `json:"status"` // สถานะ order
	Items  []UpdateOrderItemRequest `json:"items"`  // optional สำหรับแก้ไข order items
}

type UpdateOrderItemRequest struct {
	// ID        uint    `json:"id"` // ID ของ order item
	ProductID *uint   `json:"product_id,omitempty"`
	GroupID   *uint   `json:"group_id,omitempty"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"` // ราคาที่ update
}
