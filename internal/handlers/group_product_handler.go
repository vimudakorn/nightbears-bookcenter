package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vimudakorn/internal/domain"
	groupproductrequest "github.com/vimudakorn/internal/request/group_product_request"
	"github.com/vimudakorn/internal/usecases"
)

type GroupProductHandler struct {
	usecases *usecases.GroupProductUsecase
}

func NewGroupProductUsecase(uc *usecases.GroupProductUsecase) *GroupProductHandler {
	return &GroupProductHandler{usecases: uc}
}

func (h *GroupProductHandler) GetByID(c *fiber.Ctx) error {
	groupID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	isIDExist, err := h.usecases.IsGroupIDExists(uint(groupID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to check group id"})
	}
	if !isIDExist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Group ID not found"})
	}
	groupProducts, err := h.usecases.GetProductByGroupID(uint(groupID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get product in group id",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": groupProducts,
	})
}

func (h *GroupProductHandler) AddProductInGroup(c *fiber.Ctx) error {
	groupID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	isIDExist, err := h.usecases.IsGroupIDExists(uint(groupID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to check group id"})
	}
	if !isIDExist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Group ID not found"})
	}
	var req groupproductrequest.AddProductInGroupRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	if req.Quantity < 1 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Quantity must be at least 1",
		})
	}
	groupProduct := &domain.GroupProduct{
		GroupID:   uint(groupID),
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	if err := h.usecases.AddProductInGroup(groupProduct); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to add product to group",
		})
	}

	group, err := h.usecases.FindByID(uint(groupID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to find group id",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    group,
		"message": "Product added to group successfully",
	})
}

// func (h *GroupProductHandler) AddMultiProductInGroup(c *fiber.Ctx) error {
// 	groupID, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
// 	}
// 	isIDExist, err := h.usecases.IsGroupIDExists(uint(groupID))
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to check group id"})
// 	}
// 	if !isIDExist {
// 		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Group ID not found"})
// 	}
// 	var req groupproductrequest.AddMultipleProductInGroupRequest
// 	if err := c.BodyParser(&req); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
// 	}

// 	var products []domain.GroupProduct

// 	for _, product := range req.Products {
// 		if product.Quantity < 1 {
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"error": "Quantity must be at least 1",
// 			})
// 		}

// 		isExist, err := h.usecases.IsProductInGroupID(uint(groupID), product.ProductID)
// 		if err != nil {
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": "Failed to check product in group",
// 			})
// 		}
// 		if isExist {
// 			groupProductExist, err := h.usecases.FindByGroupAndProductID(uint(groupID), product.ProductID)
// 			if err != nil {
// 				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 					"error": "Failed to check product by group id and product id",
// 				})
// 			}
// 			groupProduct := domain.GroupProduct{
// 				GroupID:   uint(groupID),
// 				ProductID: product.ProductID,
// 				Quantity:  groupProductExist.Quantity + product.Quantity,
// 			}

// 			err = h.usecases.UpdateProductInGroup(&groupProduct)
// 		} else {
// 			groupProduct := domain.GroupProduct{
// 				GroupID:   uint(groupID),
// 				ProductID: product.ProductID,
// 				Quantity:  product.Quantity,
// 			}
// 			err = h.usecases.CreateProductInGroup(&groupProduct)
// 		}

// 		products = append(products, groupProduct)
// 	}
// 	group, err := h.usecases.AddMultiProductInGroup(products)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": "Failed to add products to group",
// 		})
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
// 		"data":    group,
// 		"message": "Products added to group successfully",
// 	})

// }

func (h *GroupProductHandler) AddMultiProductInGroup(c *fiber.Ctx) error {
	groupID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid group ID"})
	}

	isIDExist, err := h.usecases.IsGroupIDExists(uint(groupID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to check group ID"})
	}
	if !isIDExist {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Group ID not found"})
	}

	var req groupproductrequest.AddMultipleProductInGroupRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	for _, product := range req.Products {
		if product.Quantity < 1 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Quantity must be at least 1",
			})
		}

		isProductExist, err := h.usecases.IsProductIDExists(product.ProductID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to check if product ID exists",
			})
		}
		if !isProductExist {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": fmt.Sprintf("Product ID %d not found", product.ProductID),
			})
		}

		existingProduct, err := h.usecases.FindByGroupAndProductID(uint(groupID), product.ProductID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to check product in group",
			})
		}

		if existingProduct != nil && existingProduct.ID != 0 {
			existingProduct.Quantity += product.Quantity
			err = h.usecases.UpdateProductInGroup(existingProduct)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to update product in group",
				})
			}
		} else {
			newProduct := domain.GroupProduct{
				GroupID:   uint(groupID),
				ProductID: product.ProductID,
				Quantity:  product.Quantity,
			}
			if err := h.usecases.AddProductInGroup(&newProduct); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to create product in group",
				})
			}
		}
	}

	group, err := h.usecases.FindByID(uint(groupID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to find group by ID",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":    group,
		"message": "Products added or updated in group successfully",
	})
}
