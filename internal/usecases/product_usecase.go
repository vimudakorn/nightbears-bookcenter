package usecases

import "github.com/vimudakorn/internal/domain"

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
