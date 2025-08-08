package usecases

import (
	"fmt"

	"github.com/vimudakorn/internal/domain"
	groupproductrequest "github.com/vimudakorn/internal/request/group_product_request"
)

type GroupProductUsecase struct {
	groupProductRepo domain.GroupProductRepository
	groupRepo        domain.GroupRepository
	productRepo      domain.ProductRepository
}

func NewGroupProductUsecase(gpRepo domain.GroupProductRepository, gRepo domain.GroupRepository, pRepo domain.ProductRepository) *GroupProductUsecase {
	return &GroupProductUsecase{
		groupProductRepo: gpRepo,
		groupRepo:        gRepo,
		productRepo:      pRepo,
	}
}

func (gp *GroupProductUsecase) AddProductInGroup(groupProduct *domain.GroupProduct) error {
	return gp.groupProductRepo.Create(groupProduct)
}

func (gp *GroupProductUsecase) IsProductIDExists(productID uint) (bool, error) {
	return gp.productRepo.IsProductIDExists(productID)
}

func (gp *GroupProductUsecase) FindByID(groupID uint) (*domain.Group, error) {
	return gp.groupRepo.FindByID(groupID)
}

func (gp *GroupProductUsecase) AddMultiProductInGroup(groupProducts []domain.GroupProduct) (*domain.Group, error) {
	var groupID uint
	if len(groupProducts) > 0 {
		groupID = groupProducts[0].GroupID
	}

	err := gp.groupProductRepo.CreateMulti(groupProducts)
	if err != nil {
		return nil, err
	}

	return gp.groupRepo.FindByID(groupID)
}

func (uc *GroupProductUsecase) IsProductInGroupID(groupID, productID uint) (bool, error) {
	return uc.groupProductRepo.IsProductInGroupID(groupID, productID)
}

func (uc *GroupProductUsecase) UpdateProductInGroup(product *domain.GroupProduct) error {
	return uc.groupProductRepo.Update(product)
}

func (gp *GroupProductUsecase) IsGroupIDExists(groupID uint) (bool, error) {
	return gp.groupRepo.IsGroupIDExists(groupID)
}

func (gp *GroupProductUsecase) FindByGroupAndProductID(groupID uint, productID uint) (*domain.GroupProduct, error) {
	return gp.groupProductRepo.FindByGroupAndProductID(groupID, productID)
}

// func (gp *GroupProductUsecase) GetProductByGroupID(groupID uint) ([]domain.GroupProduct, error) {
// 	return gp.groupProductRepo.GetProductByGroupID(groupID)
// }

func (gp *GroupProductUsecase) GetProductByGroupID(groupID uint) ([]groupproductrequest.GroupProductWithDetail, error) {
	return gp.groupProductRepo.GetProductByGroupID(groupID)
}

func (gp *GroupProductUsecase) AddOrUpdateProductsInGroup(groupID uint, products []domain.GroupProduct) error {
	exists, err := gp.groupRepo.IsGroupIDExists(groupID)
	if err != nil {
		return fmt.Errorf("failed to check group: %w", err)
	}
	if !exists {
		return fmt.Errorf("group ID %d not found", groupID)
	}

	for _, p := range products {
		ok, err := gp.productRepo.IsProductIDExists(p.ProductID)
		if err != nil {
			return fmt.Errorf("failed to check product ID %d: %w", p.ProductID, err)
		}
		if !ok {
			return fmt.Errorf("product ID %d not found", p.ProductID)
		}
	}

	// ส่งเข้า repo ที่จัดการ transaction เอง
	return gp.groupProductRepo.AddOrUpdateMulti(groupID, products)
}

func (gp *GroupProductUsecase) UpdateProductsInGroup(groupID uint, updates []domain.GroupProduct) error {
	return gp.groupProductRepo.UpdateProductsInGroupID(groupID, updates)
}

func (gp *GroupProductUsecase) DeleteProductInGroup(groupID uint, productID uint) error {
	return gp.groupProductRepo.Delete(groupID, productID)
}
