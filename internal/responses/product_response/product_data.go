package productresponse

type ProductdataResponse struct {
	ID           uint    `json:"id"`
	ProductCode  int     `json:"productCode"`
	ProductType  string  `json:"productType"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float64 `json:"price"`
	SalePrice    float64 `json:"salePrice"`
	Stock        int     `json:"stock"`
	ImageURL     string  `json:"imageUrl"`
	CategoryID   *uint   `json:"categoryId"`
	ProductImage []struct {
		Image_url string `json:"image_url"`
	} `json:"productImage"`
	Tags []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"tags"`
	BookData     *BookProductDataResponse     `json:"book_data,omitempty"`
	LearningData *LearningProductDataResponse `json:"learning_data,omitempty"`
	OfficeData   *OfficeProductDataResponse   `json:"office_data,omitempty"`
}
type BookProductDataResponse struct {
	Subject          string // "รายวิชา"
	LearningArea     string // "กลุ่มสาระการเรียนรู้"
	Grade            string // "ชั้น"
	Publisher        string // "ผู้จัดพิมพ์"
	Editor           string // "ผู้เรียบเรียง"
	PublishYear      string // "ปี พ.ศ. ที่เผยแพร่"
	Size             string // "ขนาด"
	PageCount        string // "จำนวนหน้า"
	Paper            string // "กระดาษ"
	PrintType        string // "พิมพ์"
	Weight           string // "น้ำหนัก"
	LicenseURL       string // "ใบอนุญาต"
	CertificateURL   string // "ใบประกาศ"
	WarrantyURL      string // "ใบประกัน"
	SampleContentURL string // "ตัวอย่างเนื้อหา"
	Author           string // "ผู้เรียบเรียง" (or keep Author separate if needed)
	ISBN             string // if available
}

type LearningProductDataResponse struct {
	Brand    string
	Material string
}

type OfficeProductDataResponse struct {
	Color string
	Size  string
}
