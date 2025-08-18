package productresponse

import "github.com/vimudakorn/internal/domain"

func MapProductToResponse(p domain.Product) ProductdataResponse {
	resp := ProductdataResponse{
		ID:          p.ID,
		ProductCode: p.ProductCode,
		ProductType: p.ProductType,
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		SalePrice:   p.Price - p.Discount,
		Stock:       p.Stock,
		ImageURL:    p.ImageURL,
		CategoryID:  p.CategoryID,
	}

	for _, img := range p.ProductImages {
		resp.ProductImage = append(resp.ProductImage, struct {
			Image_url string `json:"image_url"`
		}{Image_url: img.ImageURL})
	}

	for _, tag := range p.Tags {
		resp.Tags = append(resp.Tags, struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
		}{ID: tag.ID, Name: tag.Name})
	}

	// ถ้าเป็น book → เติมข้อมูล
	if p.ProductType == "book" && p.Book != nil {
		resp.BookData = &BookProductDataResponse{
			Subject:          p.Book.Subject,
			LearningArea:     p.Book.LearningArea,
			Grade:            p.Book.Grade,
			Publisher:        p.Book.Publisher,
			Editor:           p.Book.Editor,
			PublishYear:      p.Book.PublishYear,
			Size:             p.Book.Size,
			PageCount:        p.Book.PageCount,
			Paper:            p.Book.Paper,
			PrintType:        p.Book.PrintType,
			Weight:           p.Book.Weight,
			LicenseURL:       p.Book.LicenseURL,
			CertificateURL:   p.Book.CertificateURL,
			WarrantyURL:      p.Book.WarrantyURL,
			SampleContentURL: p.Book.SampleContentURL,
			Author:           p.Book.Author,
			ISBN:             p.Book.ISBN,
		}
	}

	if p.ProductType == "learning" && p.LearningSupply != nil {
		resp.LearningData = &LearningProductDataResponse{
			Material: p.LearningSupply.Material,
			Brand:    p.LearningSupply.Brand,
		}
	}

	return resp
}
