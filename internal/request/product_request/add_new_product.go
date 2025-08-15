package productrequest

import (
	"github.com/vimudakorn/internal/domain"
)

type AddNewBookRequest struct {
	ProductCode      int                   `json:"product_code" validate:"unique, required"`
	ProductType      string                `json:"product_type"`
	Name             string                `json:"name"`
	Description      string                `json:"description"`
	Price            float64               `json:"price" validate:"gt=0, required"`
	SalePrice        float64               `json:"sale_price" validate:"gt=0, required"`
	Stock            int                   `json:"stock"`
	ImageURL         string                `json:"image_url"`
	CategoryID       *uint                 `json:"category_id"`
	ProductImage     []domain.ProductImage `json:"product_image"`
	TagIDs           []uint                `json:"tag_ids,omitempty"`
	Subject          string                `json:"subject"`            // "รายวิชา"
	LearningArea     string                `json:"learning_area"`      // "กลุ่มสาระการเรียนรู้"
	Grade            string                `json:"grade"`              // "ชั้น"
	Publisher        string                `json:"publisher"`          // "ผู้จัดพิมพ์"
	Editor           string                `json:"editor"`             // "ผู้เรียบเรียง"
	PublishYear      string                `json:"publish_year"`       // "ปี พ.ศ. ที่เผยแพร่"
	Size             string                `json:"size"`               // "ขนาด"
	PageCount        string                `json:"page_count"`         // "จำนวนหน้า"
	Paper            string                `json:"paper"`              // "กระดาษ"
	PrintType        string                `json:"print_type"`         // "พิมพ์"
	Weight           string                `json:"weight"`             // "น้ำหนัก"
	LicenseURL       string                `json:"license_url"`        // "ใบอนุญาต"
	CertificateURL   string                `json:"certificate_url"`    // "ใบประกาศ"
	WarrantyURL      string                `json:"warranty_url"`       // "ใบประกัน"
	SampleContentURL string                `json:"sample_content_url"` // "ตัวอย่างเนื้อหา"
	Author           string                `json:"author"`             // "ผู้เรียบเรียง" (or keep Author separate if needed)
	ISBN             string                `json:"isbn"`               // if available
}

type AddNewLearningRequest struct {
	ProductCode  int                   `json:"product_code" validate:"unique, required"`
	ProductType  string                `json:"product_type"`
	Name         string                `json:"name"`
	Description  string                `json:"description"`
	Price        float64               `json:"price"`
	SalePrice    float64               `json:"sale_price" validate:"gt=0, required"`
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
	SalePrice    float64               `json:"sale_price" validate:"gt=0, required"`
	Stock        int                   `json:"stock"`
	ImageURL     string                `json:"image_url"`
	CategoryID   *uint                 `json:"category_id"`
	ProductImage []domain.ProductImage `json:"product_image"`
	TagIDs       []uint                `json:"tag_ids,omitempty"`
	Color        string                `json:"color"`
	Size         string                `json:"size"`
}
