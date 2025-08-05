package handlers

import "github.com/vimudakorn/internal/usecases"

type TagHandler struct {
	usecases *usecases.TagUsecase
}

func NewTagHandler(uc *usecases.TagUsecase) *TagHandler {
	return &TagHandler{usecases: uc}
}
