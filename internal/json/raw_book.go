package jsons

type RawBook struct {
	Type             string  `json:"ประเภท"`
	Name             string  `json:"ชื่อหนังสือ"`
	Subject          string  `json:"รายวิชา"`
	LearningArea     string  `json:"กลุ่มสาระการเรียนรู้"`
	Grade            string  `json:"ชั้น"`
	Publisher        string  `json:"ผู้จัดพิมพ์"`
	Editor           string  `json:"ผู้เรียบเรียง"`
	PublishYear      string  `json:"ปี พ.ศ. ที่เผยแพร่"`
	Size             string  `json:"ขนาด"`
	PageCount        string  `json:"จำนวนหน้า"`
	Paper            string  `json:"กระดาษ"`
	PrintType        string  `json:"พิมพ์"`
	Weight           string  `json:"น้ำหนัก"`
	Price            string  `json:"ราคา"`
	ImageURL         string  `json:"ภาพปก"`
	CertificateURL   *string `json:"ใบประกาศ"`
	LicenseURL       *string `json:"ใบอนุญาต"`
	WarrantyURL      *string `json:"ใบประกัน"`
	SampleContentURL *string `json:"ตัวอย่างเนื้อหา"`
}
