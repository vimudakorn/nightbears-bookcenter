package productrequest

import (
	"github.com/vimudakorn/internal/domain"
)

type AddNewBookRequest struct {
	ProductCode  int                   `json:"product_code" validate:"unique, required"`
	ProductType  string                `json:"product_type"`
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	Price        float64               `json:"price"`
	Stock        int                   `json:"stock"`
	ImageURL     string                `json:"image_url"`
	CategoryID   *uint                 `json:"category_id"`
	ProductImage []domain.ProductImage `json:"product_image"`
	TagIDs       []uint                `json:"tag_ids,omitempty"`
	Author       string                `json:"author"`
	ISBN         string                `json:"isbn"`
}

type AddNewLearningRequest struct {
	ProductCode  int                   `json:"product_code" validate:"unique, required"`
	ProductType  string                `json:"product_type"`
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	Price        float64               `json:"price"`
	Stock        int                   `json:"stock"`
	ImageURL     string                `json:"image_url"`
	CategoryID   *uint                 `json:"category_id"`
	ProductImage []domain.ProductImage `json:"product_image"`
	TagIDs       []uint                `json:"tag_ids,omitempty"`
	Brand        string                `json:"brand"`
	Material     string                `json:"material"`
}

type AddNewOfficeRequest struct {
	ProductCode  int                   `json:"product_code" validate:"unique, required"`
	ProductType  string                `json:"product_type"`
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	Price        float64               `json:"price"`
	Stock        int                   `json:"stock"`
	ImageURL     string                `json:"image_url"`
	CategoryID   *uint                 `json:"category_id"`
	ProductImage []domain.ProductImage `json:"product_image"`
	TagIDs       []uint                `json:"tag_ids,omitempty"`
	Color        string                `json:"color"`
	Size         string                `json:"size"`
}
