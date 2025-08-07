package usecases

import "github.com/vimudakorn/internal/domain"

type GroupProductUsecase struct {
	groupProductRepo domain.GroupProductRepository
	groupRepo        domain.GroupRepository
}

func NewGroupProductUsecase(gpRepo domain.GroupProductRepository, gRepo domain.GroupRepository) *GroupProductUsecase {
	return &GroupProductUsecase{groupProductRepo: gpRepo}
}

func (gp *GroupProductUsecase) AddProductInGroup(groupProduct *domain.GroupProduct) (*domain.Group, error) {
	err := gp.groupProductRepo.Create(groupProduct)
	if err != nil {
		return nil, err
	}

	group, err := gp.groupRepo.FindByID(groupProduct.GroupID)
	if err != nil {
		return nil, err
	}
	return group, nil
}

func (gp *GroupProductUsecase) IsGroupIDExists(groupID uint) (bool, error) {
	return gp.groupRepo.IsGroupIDExists(groupID)
}
