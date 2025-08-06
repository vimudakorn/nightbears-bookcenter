package handlers

import "github.com/vimudakorn/internal/usecases"

type GroupProductHandler struct {
	usecases *usecases.GroupProductUsecase
}

func NewGroupProductUsecase(uc *usecases.GroupProductUsecase) *GroupProductHandler {
	return &GroupProductHandler{usecases: uc}
}
