package usecases

import "github.com/vimudakorn/internal/domain"

type ProductUsecase struct {
	repo domain.ProductRepository
}

func NewProductUsecase(r domain.ProductRepository) *ProductUsecase {
	return &ProductUsecase{repo: r}
}

func (r *ProductUsecase) GetPagination(page int, limit int, search string, sortBy string, sortOrder string) ([]domain.Product, int64, error) {
	return r.repo.GetPagination(page, limit, search, sortBy, sortOrder)
}

func (r *ProductUsecase) AddNewProduct(product *domain.Product) error {
	return r.repo.CreateProduct(product)
}

func (r *ProductUsecase) CreateBook(book *domain.Book) error {
	return r.repo.CreateBook(book)
}

func (r *ProductUsecase) CreateLearning(learning *domain.LearningSupply) error {
	return r.repo.CreateLearning(learning)
}

func (r *ProductUsecase) CreateOffice(office *domain.OfficeSupply) error {
	return r.repo.CreateOffice(office)
}

func (r *ProductUsecase) FindByID(productID uint) (*domain.Product, error) {
	return r.repo.FindByID(productID)
}

func (r *ProductUsecase) Update(product *domain.Product) error {
	return r.repo.Update(product)
}

func (r *ProductUsecase) Delete(productID uint) error {
	return r.repo.Delete(productID)
}

func (r *ProductUsecase) DeleteBook(bookId uint) error {
	return r.repo.DeleteBook(bookId)
}

func (r *ProductUsecase) DeleteOffice(officeId uint) error {
	return r.repo.DeleteOffice(officeId)
}

func (r *ProductUsecase) DeleteLearning(learningId uint) error {
	return r.repo.DeleteLearning(learningId)
}
