package usecases

import (
	"github.com/vimudakorn/internal/domain"
	errorsres "github.com/vimudakorn/internal/domain/errors_res"
	grouprequest "github.com/vimudakorn/internal/request/group_request"
	"gorm.io/gorm"
)

type GroupUsecase struct {
	groupRepo        domain.GroupRepository
	groupProductRepo domain.GroupProductRepository
}

func NewGroupUsecase(groupRepo domain.GroupRepository, groupProductRepo domain.GroupProductRepository) *GroupUsecase {
	return &GroupUsecase{groupRepo: groupRepo, groupProductRepo: groupProductRepo}
}

func (r *GroupUsecase) GetPagination(page int, limit int, search string, sortBy string, orderBy string) ([]domain.Group, int64, error) {
	return r.groupRepo.GetPagination(page, limit, search, sortBy, orderBy)
}

func (r *GroupUsecase) Create(group *domain.Group) error {
	return r.groupRepo.Create(group)
}

func (r *GroupUsecase) IsNameAndEduExist(name string, level string) (bool, error) {
	return r.groupRepo.IsNameAndEduExist(name, level)
}

func (r *GroupUsecase) FindByID(id uint) (*domain.Group, error) {
	return r.groupRepo.FindByID(id)
}

func (g *GroupUsecase) UpdateGroup(id uint, req grouprequest.UpdateGroupRequest) error {
	group, err := g.groupRepo.FindByID(id)
	if err != nil {
		return errorsres.ErrGroupNameExist
	}

	if req.Name != "" && req.EduLevel != "" {
		isExist, err := g.groupRepo.IsNameAndEduExist(req.Name, req.EduLevel)
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

	return g.groupRepo.Update(group)
}

func (g *GroupUsecase) Delete(id uint) error {
	return g.groupRepo.Delete(id)
}

func (g *GroupUsecase) CreateGroupWithProducts(req grouprequest.CreateGroupWithProductsRequest) error {
	// เช็คชื่อกลุ่มซ้ำ
	isExist, err := g.groupRepo.IsNameAndEduExist(req.Name, req.EduLevel)
	if err != nil {
		return err
	}
	if isExist {
		return errorsres.ErrGroupNameExist
	}

	if req.SalePrice <= 0 {
		return errorsres.ErrInvalidSalePrice
	}

	// ทำ Transaction
	return g.groupRepo.GetDB().Transaction(func(tx *gorm.DB) error {
		group := domain.Group{
			Name:        req.Name,
			EduLevel:    req.EduLevel,
			Description: req.Description,
			SalePrice:   req.SalePrice,
		}

		if err := g.groupRepo.CreateWithProduct(tx, &group); err != nil {
			return err
		}

		for _, p := range req.Products {
			if p.Quantity <= 0 {
				continue // หรือ return error ก็ได้
			}
			pg := domain.GroupProduct{
				GroupID:   group.ID,
				ProductID: p.ProductID,
				Quantity:  p.Quantity,
			}
			if err := g.groupProductRepo.CreateWithProduct(tx, &pg); err != nil {
				return err
			}
		}

		return nil
	})
}

func (g *GroupUsecase) IsGroupIDExists(groupID uint) (bool, error) {
	return g.groupRepo.IsGroupIDExists(groupID)
}
