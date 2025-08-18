package usecases

import (
	"github.com/vimudakorn/internal/domain"
	"github.com/vimudakorn/internal/utils"
)

type ProductUsecase struct {
	productRepo domain.ProductRepository
	tagRepo     domain.TagRepository
}

func NewProductUsecase(productRepo domain.ProductRepository, tagRepo domain.TagRepository) *ProductUsecase {
	return &ProductUsecase{productRepo: productRepo, tagRepo: tagRepo}
}

func (r *ProductUsecase) GetPagination(page int, limit int, search string, sortBy string, sortOrder string) ([]domain.Product, int64, error) {
	return r.productRepo.GetPagination(page, limit, search, sortBy, sortOrder)
}

func (r *ProductUsecase) AddNewProduct(product *domain.Product) error {
	return r.productRepo.CreateProduct(product)
}

func (r *ProductUsecase) CreateBook(book *domain.Book) error {
	return r.productRepo.CreateBook(book)
}

func (r *ProductUsecase) CreateLearning(learning *domain.LearningSupply) error {
	return r.productRepo.CreateLearning(learning)
}

func (r *ProductUsecase) CreateOffice(office *domain.OfficeSupply) error {
	return r.productRepo.CreateOffice(office)
}

func (r *ProductUsecase) FindByID(productID uint) (*domain.Product, error) {
	return r.productRepo.FindByID(productID)
}

func (r *ProductUsecase) Update(product *domain.Product) error {
	return r.productRepo.Update(product)
}

func (r *ProductUsecase) Delete(productID uint) error {
	return r.productRepo.Delete(productID)
}

func (r *ProductUsecase) DeleteBook(bookId uint) error {
	return r.productRepo.DeleteBook(bookId)
}

func (r *ProductUsecase) DeleteOffice(officeId uint) error {
	return r.productRepo.DeleteOffice(officeId)
}

func (r *ProductUsecase) DeleteLearning(learningId uint) error {
	return r.productRepo.DeleteLearning(learningId)
}

func (r *ProductUsecase) FindBookID(productID uint) (uint, error) {
	return r.productRepo.FindBookID(productID)
}

func (r *ProductUsecase) FindLearningID(productID uint) (uint, error) {
	return r.productRepo.FindLearningID(productID)
}

func (r *ProductUsecase) FindOfficeID(productID uint) (uint, error) {
	return r.productRepo.FindOfficeID(productID)
}

func (r *ProductUsecase) GetTagsByIDs(ids []uint) ([]domain.Tag, error) {
	return r.tagRepo.GetTagsByIDs(ids)
}

func (r *ProductUsecase) AddTagsToProduct(productID uint, tagIDs []uint) error {
	product, err := r.productRepo.FindByID(productID)
	if err != nil {
		return err
	}

	tags, err := r.tagRepo.GetTagsByIDs(tagIDs)
	if err != nil {
		return err
	}

	product.Tags = tags

	if err := r.productRepo.Update(product); err != nil {
		return err
	}

	return nil
}

func (u *ProductUsecase) CreateFromJSON(path string) error {
	products, err := utils.ReadProductsFromJSON(path)
	if err != nil {
		return err
	}

	for _, p := range products {
		if err := u.productRepo.CreateProduct(&p); err != nil {
			return err
		}
	}
	return nil
}

// func (u *ProductUsecase) ImportBooksFromJSON(filename string, categoryID uint) error {
// 	file, err := os.Open(filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	var raws []jsons.RawBook
// 	if err := json.NewDecoder(file).Decode(&raws); err != nil {
// 		return err
// 	}

// 	for _, r := range raws {
// 		price := utils.ParsePrice(r.Price)

// 		p := domain.Product{
// 			ProductCode: 0,
// 			ProductType: "book",
// 			Name:        r.Name,
// 			Price:       price,
// 			ImageURL:    r.ImageURL,
// 			CategoryID:  &categoryID,
// 			ProductImages: []domain.ProductImage{
// 				{ImageURL: r.ImageURL},
// 			},
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
// 				LicenseURL:       utils.GetString(r.LicenseURL),
// 				CertificateURL:   r.CertificateURL,
// 				WarrantyURL:      utils.GetString(r.WarrantyURL),
// 				SampleContentURL: utils.GetString(r.SampleContentURL),
// 				Author:           r.Editor,
// 			},
// 		}

// 		if err := u.productRepo.CreateProduct(&p); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
