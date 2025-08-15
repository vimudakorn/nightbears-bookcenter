package domain

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	ProductID uint `gorm:"uniqueIndex"` // One-to-one

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
	// CoverImageURL    string // "ภาพปก"
}
