package usecases

import "github.com/vimudakorn/internal/domain"

type TagUsecase struct {
	repo domain.TagRepository
}

func NewTagUsecase(r domain.TagRepository) *TagUsecase {
	return &TagUsecase{repo: r}
}
