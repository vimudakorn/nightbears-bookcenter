package responses

import "github.com/vimudakorn/internal/domain"

type ProductdataResponse struct {
	ID           uint                  `json:"id"`
	ProductCode  int                   `json:"product_code"`
	ProductType  string                `json:"product_type"`
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	Price        float64               `json:"price"`
	Stock        int                   `json:"stock"`
	ImageURL     string                `json:"image_url"`
	CategoryID   uint                  `json:"category_id"`
	ProductImage []domain.ProductImage `json:"product_image"`
}
