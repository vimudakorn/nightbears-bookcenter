package utils

import (
	"encoding/json"
	"os"

	"github.com/vimudakorn/internal/domain"
)

// อ่าน JSON แล้ว map ไปยัง domain.Product + domain.Book
// func ReadProductsFromJSON(path string) ([]domain.Product, error) {
// 	file, err := os.ReadFile(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var raws []jsons.RawBook
// 	if err := json.Unmarshal(file, &raws); err != nil {
// 		return nil, err
// 	}

// 	var products []domain.Product
// 	for _, r := range raws {
// 		// price := ParsePrice(r.Price)

// 		p := domain.Product{
// 			ProductType: "book",
// 			Name:        r.Name,
// 			Price:       ParsePrice(r.Price),
// 			ImageURL:    r.ImageURL,
// 			Book: &domain.Book{
// 				Subject:          r.Subject,
// 				LearningArea:     r.LearningArea,
// 				Grade:            r.Grade,
// 				Publisher:        r.Publisher,
// 				Editor:           r.Editor,
// 				PublishYear:      r.PublishYear,
// 				Size:             r.Size,
// 				PageCount:        r.PageCount,
// 				Paper:            r.Paper,
// 				PrintType:        r.PrintType,
// 				Weight:           r.Weight,
// 				CertificateURL:   safeString(r.CertificateURL),
// 				LicenseURL:       safeString(r.LicenseURL),
// 				WarrantyURL:      safeString(r.WarrantyURL),
// 				SampleContentURL: safeString(r.SampleContentURL),
// 				Author:           r.Editor,
// 			},
// 			ProductImages: []domain.ProductImage{
// 				{ImageURL: r.ImageURL},
// 			},
// 		}

// 		products = append(products, p)
// 	}

// 	return products, nil
// }

// ฟังก์ชันช่วยแปลง pointer -> string
func safeString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func ReadProductsFromJSON(path string) ([]domain.Product, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var raw []map[string]interface{}
	if err := json.Unmarshal(file, &raw); err != nil {
		return nil, err
	}

	var products []domain.Product
	for _, r := range raw {
		// map fields จาก JSON → domain.Product
		// ตัวอย่าง
		p := domain.Product{
			ProductType: "book",
			Name:        r["ชื่อหนังสือ"].(string),
			Price:       ParsePrice(r["ราคา"].(string)),
			ImageURL:    r["ภาพปก"].(string),
			Book: &domain.Book{
				Subject:          r["รายวิชา"].(string),
				LearningArea:     r["กลุ่มสาระการเรียนรู้"].(string),
				Grade:            r["ชั้น"].(string),
				Publisher:        r["ผู้จัดพิมพ์"].(string),
				Editor:           r["ผู้เรียบเรียง"].(string),
				PublishYear:      r["ปี พ.ศ. ที่เผยแพร่"].(string),
				Size:             r["ขนาด"].(string),
				PageCount:        r["จำนวนหน้า"].(string),
				Paper:            r["กระดาษ"].(string),
				PrintType:        r["พิมพ์"].(string),
				Weight:           r["น้ำหนัก"].(string),
				CertificateURL:   safeString(ptrString(r["ใบประกาศ"])),
				LicenseURL:       safeString(ptrString(r["ใบอนุญาต"])),
				WarrantyURL:      safeString(ptrString(r["ใบประกัน"])),
				SampleContentURL: safeString(ptrString(r["ตัวอย่างเนื้อหา"])),
			}}
		products = append(products, p)
	}

	return products, nil
}

func ptrString(s interface{}) *string {
	if s == nil {
		return nil
	}
	str, ok := s.(string)
	if !ok {
		return nil
	}
	return &str
}
