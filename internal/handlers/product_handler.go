package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vimudakorn/constants"
	"github.com/vimudakorn/internal/domain"
	productrequest "github.com/vimudakorn/internal/request/product_request"
	"github.com/vimudakorn/internal/responses"
	"github.com/vimudakorn/internal/usecases"
)

type ProductHandler struct {
	usecases *usecases.ProductUsecase
}

func NewProductHandler(uc *usecases.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecases: uc}
}

func (h *ProductHandler) GetAll(c *fiber.Ctx) error {
	name := c.Query("name")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	sortBy := c.Query("sortBy", "name")
	sortOrder := c.Query("sortOrder", "asc")
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}

	products, count, err := h.usecases.GetPagination(page, limit, name, sortBy, sortOrder)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	var res []responses.ProductdataResponse
	for _, product := range products {
		res = append(res, responses.ProductdataResponse{
			ID:           product.ID,
			ProductCode:  product.ProductCode,
			ProductType:  product.ProductType,
			Name:         product.Name,
			Description:  product.Description,
			Price:        product.Price,
			Stock:        product.Stock,
			ImageURL:     product.ImageURL,
			CategoryID:   *product.CategoryID,
			ProductImage: product.ProductImages,
			TagIDs:       product.Tags,
		})
	}
	return c.JSON(fiber.Map{
		"full":       products,
		"data":       res,
		"page":       page,
		"limit":      limit,
		"count":      count,
		"totalPages": (int(count) + limit - 1) / limit,
	})
}

func (h *ProductHandler) AddNewProduct(c *fiber.Ctx) error {
	productType := c.Query("type")

	switch productType {
	case "book":
		var req productrequest.AddNewBookRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		var tags []domain.Tag
		if len(req.TagIDs) > 0 {
			var err error
			tags, err = h.usecases.GetTagsByIDs(req.TagIDs)
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "failed to fetch tags")
			}
		}

		if req.SalePrice > req.Price {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to set sale price more expansive than price"})
		}

		product := &domain.Product{
			ProductCode:   req.ProductCode,
			ProductType:   req.ProductType,
			Name:          req.Name,
			Description:   req.Description,
			Price:         req.Price,
			Discount:      req.Price - req.SalePrice,
			Stock:         req.Stock,
			ImageURL:      req.ImageURL,
			CategoryID:    req.CategoryID,
			ProductImages: req.ProductImage,
			Tags:          tags,
		}

		if err := h.usecases.AddNewProduct(product); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to create product")
		}

		book := &domain.Book{
			ProductID:        product.ID,
			Subject:          req.Subject,
			LearningArea:     req.LearningArea,
			Grade:            req.Grade,
			Publisher:        req.Grade,
			Editor:           req.Editor,
			PublishYear:      req.PublishYear,
			Size:             req.Size,
			PageCount:        req.PageCount,
			Paper:            req.Paper,
			PrintType:        req.PrintType,
			Weight:           req.Weight,
			LicenseURL:       req.LicenseURL,
			CertificateURL:   req.CertificateURL,
			WarrantyURL:      req.WarrantyURL,
			SampleContentURL: req.SampleContentURL,
			Author:           req.Author,
			ISBN:             req.ISBN,
		}

		if err := h.usecases.CreateBook(book); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to create book")
		}
	case "learning":
		var req productrequest.AddNewLearningRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		var tags []domain.Tag
		if len(req.TagIDs) > 0 {
			var err error
			tags, err = h.usecases.GetTagsByIDs(req.TagIDs)
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "failed to fetch tags")
			}
		}

		if req.SalePrice > req.Price {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to set sale price more expansive than price"})
		}

		product := &domain.Product{
			ProductCode:   req.ProductCode,
			ProductType:   req.ProductType,
			Name:          req.Name,
			Description:   req.Description,
			Price:         req.Price,
			Discount:      req.Price - req.SalePrice,
			Stock:         req.Stock,
			ImageURL:      req.ImageURL,
			CategoryID:    req.CategoryID,
			ProductImages: req.ProductImage,
			Tags:          tags,
		}

		if err := h.usecases.AddNewProduct(product); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to create product")
		}

		learning := &domain.LearningSupply{
			ProductID: product.ID,
			Brand:     req.Brand,
			Material:  req.Material,
		}

		if err := h.usecases.CreateLearning(learning); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to create learning supply")
		}
	case "office":
		var req productrequest.AddNewOfficeRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		var tags []domain.Tag
		if len(req.TagIDs) > 0 {
			var err error
			tags, err = h.usecases.GetTagsByIDs(req.TagIDs)
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "failed to fetch tags")
			}
		}

		if req.SalePrice > req.Price {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to set sale price more expansive than price"})
		}

		product := &domain.Product{
			ProductCode:   req.ProductCode,
			ProductType:   req.ProductType,
			Name:          req.Name,
			Description:   req.Description,
			Price:         req.Price,
			Discount:      req.Price - req.SalePrice,
			Stock:         req.Stock,
			ImageURL:      req.ImageURL,
			CategoryID:    req.CategoryID,
			ProductImages: req.ProductImage,
			Tags:          tags,
		}

		if err := h.usecases.AddNewProduct(product); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to create product")
		}

		office := &domain.OfficeSupply{
			ProductID: product.ID,
			Color:     req.Color,
			Size:      req.Size,
		}

		if err := h.usecases.CreateOffice(office); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to create office supply")
		}
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product type"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Product created successfully"})
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	var req productrequest.UpdateProduct
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid product id"})
	}

	product, err := h.usecases.FindByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
	}

	var tags []domain.Tag
	if len(req.TagIDs) > 0 {
		var err error
		tags, err = h.usecases.GetTagsByIDs(req.TagIDs)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to fetch tags")
		}
	}
	if req.SalePrice > req.Price {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to set sale price more expansive than price"})
	}

	product.Name = req.Name
	product.Price = req.Price
	product.Discount = req.Price - req.SalePrice
	product.Stock = req.Stock
	product.Tags = tags

	if err := h.usecases.Update(product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to update product"})
	}

	return c.JSON(fiber.Map{
		"message": "product updated successfully",
		"product": product,
	})
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	productID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid product id"})
	}

	requestingUserRole := c.Locals("role").(string)

	if requestingUserRole != constants.ADMIN {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "you are not authorized to delete this product",
		})
	}
	product, err := h.usecases.FindByID(uint(productID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "product not found"})
	}

	switch product.ProductType {
	case "book":
		bookID, err := h.usecases.FindBookID(product.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to find book"})
		}

		err = h.usecases.DeleteBook(bookID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete book info"})
		}

	case "learning":
		learningID, err := h.usecases.FindLearningID(product.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to find learning supply"})
		}

		err = h.usecases.DeleteLearning(learningID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete learning supply info"})
		}
	case "office":
		officeID, err := h.usecases.FindOfficeID(product.ID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to find office supply"})
		}

		err = h.usecases.DeleteOffice(officeID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to delete office supply info"})
		}
	}

	if err := h.usecases.Delete(uint(productID)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
