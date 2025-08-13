package cartitemrequest

import "fmt"

type UpdateMultiCartItemsRequest struct {
	Items []AddCartItemRequest `json:"items"`
}

type UpdateCartItemRequest struct {
	ProductID *uint `json:"product_id"`
	GroupID   *uint `json:"group_id"`
	Quantity  int   `json:"quantity"`
}

func (r *UpdateCartItemRequest) Validate() error {
	if r.ProductID == nil && r.GroupID == nil {
		return fmt.Errorf("must specify either product_id or group_id")
	}
	if r.ProductID != nil && r.GroupID != nil {
		return fmt.Errorf("only one of product_id or group_id can be specified")
	}
	if r.Quantity < 1 {
		return fmt.Errorf("quantity must be at least 1")
	}
	return nil
}
