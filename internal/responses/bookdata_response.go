package responses

import "github.com/vimudakorn/internal/domain"

type BookdataResponse struct {
	Title        string            `json:"title"`
	Author       string            `json:"author"`
	ISBN         string            `json:"isbn"`
	Quantity     uint              `json:"quantity"`
	Description  string            `json:"description"`
	Price        float64           `json:"price"`
	Discount     float64           `json:"discount"`
	DiscountType string            `json:"discount_type"`
	ImageURL     string            `json:"image_url"`
	Categories   []domain.Category `json:"categories"`
}
