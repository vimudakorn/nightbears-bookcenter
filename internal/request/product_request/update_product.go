package productrequest

type UpdateProduct struct {
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	SalePrice float64 `json:"sale_price"`
	Stock     int     `json:"stock"`
	TagIDs    []uint  `json:"tag_ids,omitempty"`
}
