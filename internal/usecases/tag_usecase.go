package usecases

import "github.com/vimudakorn/internal/domain"

type TagUsecase struct {
	repo domain.TagRepository
}

func NewTagUsecase(r domain.TagRepository) *TagUsecase {
	return &TagUsecase{repo: r}
}

func (r *TagUsecase) Create(tag *domain.Tag) error {
	return r.repo.Create(tag)
}

func (r *TagUsecase) GetAll() ([]domain.Tag, error) {
	return r.repo.FindAll()
}

func (r *TagUsecase) Update(tag *domain.Tag) error {
	return r.repo.Update(tag)
}

func (r *TagUsecase) Delete(id uint) error {
	return r.repo.Delete(id)
}

func (r *TagUsecase) GetPagination(page int, limit int, search string, sortBy string, orderBy string) ([]domain.Tag, int64, error) {
	return r.repo.GetPagination(page, limit, search, sortBy, orderBy)
}

func (r *TagUsecase) RenameTag(tagID uint, newName string) error {
	return r.repo.RenameTag(tagID, newName)
}

func (r *TagUsecase) IsTagHasUsed(id uint) (bool, error) {
	return r.repo.IsTagHasUsed(id)
}

func (r *TagUsecase) IsNameExists(name string) (bool, error) {
	return r.repo.IsNameExists(name)
}
