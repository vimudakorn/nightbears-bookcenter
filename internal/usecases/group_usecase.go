package usecases

import (
	"github.com/vimudakorn/internal/domain"
	errorsres "github.com/vimudakorn/internal/domain/errors_res"
	grouprequest "github.com/vimudakorn/internal/request/group_request"
)

type GroupUsecase struct {
	repo domain.GroupRepository
}

func NewGroupUsecase(r domain.GroupRepository) *GroupUsecase {
	return &GroupUsecase{repo: r}
}

func (r *GroupUsecase) GetPagination(page int, limit int, search string, sortBy string, orderBy string) ([]domain.Group, int64, error) {
	return r.repo.GetPagination(page, limit, search, sortBy, orderBy)
}

func (r *GroupUsecase) Create(group *domain.Group) error {
	return r.repo.Create(group)
}

func (r *GroupUsecase) IsNameAndEduExist(name string, level string) (bool, error) {
	return r.repo.IsNameAndEduExist(name, level)
}

func (r *GroupUsecase) FindByID(id uint) (*domain.Group, error) {
	return r.repo.FindByID(id)
}

func (g *GroupUsecase) UpdateGroup(id uint, req grouprequest.UpdateGroupRequest) error {
	group, err := g.repo.FindByID(id)
	if err != nil {
		return errorsres.ErrGroupNameExist
	}

	if req.Name != "" && req.EduLevel != "" {
		isExist, err := g.repo.IsNameAndEduExist(req.Name, req.EduLevel)
		if err != nil {
			return err
		}
		if isExist {
			return errorsres.ErrGroupNameExist
		}
		group.Name = req.Name
		group.EduLevel = req.EduLevel
	}

	if req.Description != "" {
		group.Description = req.Description
	}

	if req.SalePrice <= 0 {
		return errorsres.ErrInvalidSalePrice
	}
	group.SalePrice = req.SalePrice

	return g.repo.Update(group)
}

func (g *GroupUsecase) Delete(id uint) error {
	return g.repo.Delete(id)
}
