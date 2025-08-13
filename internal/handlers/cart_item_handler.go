package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/vimudakorn/internal/domain"
	cartitemrequest "github.com/vimudakorn/internal/request/cart_item_request"
	"github.com/vimudakorn/internal/usecases"
)

type CartItemHandler struct {
	usecases *usecases.CartItemUsecase
}

func NewCartItemHandler(uc *usecases.CartItemUsecase) *CartItemHandler {
	return &CartItemHandler{usecases: uc}
}

func (h *CartItemHandler) GetOwnItemInCard(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id type")
	}
	cart, err := h.usecases.GetCartByUserID(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to get cart")
	}

	cartItems, err := h.usecases.GetItemsByCartID(cart.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to check items in cart ID"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data": cartItems,
	})
}

func (h *CartItemHandler) AddOrUpdateMultiCartItems(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id type")
	}

	cart, err := h.usecases.GetCartByUserID(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to get cart")
	}

	var req cartitemrequest.AddMultiCartItemsRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	var cartItems []domain.CartItem
	for _, item := range req.Items {
		if err := item.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		if item.Quantity < 1 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Quantity must be at least 1",
			})
		}

		if item.ProductID != nil {
			exists, err := h.usecases.IsProductIDExists(*item.ProductID)
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "failed to check product ID")
			}
			if !exists {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": fmt.Sprintf("product ID %d not found", *item.ProductID)})
			}
		}

		if item.GroupID != nil {
			exists, err := h.usecases.IsGroupIDExist(*item.GroupID)
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "failed to check group ID")
			}
			if !exists {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": fmt.Sprintf("group ID %d not found", *item.GroupID)})
			}
		}

		cartItems = append(cartItems, domain.CartItem{
			CartID:    cart.ID,
			ProductID: item.ProductID,
			GroupID:   item.GroupID,
			Quantity:  item.Quantity,
		})
	}

	if err := h.usecases.AddOrUpdateMultiCartItems(cart.ID, cartItems); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to add/update cart items")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Cart items added/updated successfully",
	})
}

func (h *CartItemHandler) AddOrUpdateCartItem(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id type")
	}

	cart, err := h.usecases.GetCartByUserID(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to get cart")
	}

	var req cartitemrequest.AddCartItemRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	if err := req.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if req.ProductID != nil {
		isProductExist, err := h.usecases.IsProductIDExists(*req.ProductID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to check product")
		}
		if !isProductExist {
			return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Product ID %d not found", *req.ProductID))
		}
	}

	if req.GroupID != nil {
		isGroupExist, err := h.usecases.IsGroupIDExist(*req.GroupID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to check group")
		}
		if !isGroupExist {
			return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Group ID %d not found", *req.GroupID))
		}
	}

	item := domain.CartItem{
		CartID:    cart.ID,
		ProductID: req.ProductID,
		GroupID:   req.GroupID,
		Quantity:  req.Quantity,
	}

	if err := h.usecases.AddOrUpdateCartItem(cart.ID, &item); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to update cart item")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *CartItemHandler) UpdateItemsInCart(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id type")
	}
	cart, err := h.usecases.GetCartByUserID(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to get cart")
	}

	var req cartitemrequest.UpdateMultiCartItemsRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	var items []domain.CartItem
	for _, item := range req.Items {
		if err := item.Validate(); err != nil {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}

		if item.ProductID != nil {
			isProductExist, err := h.usecases.IsProductIDExists(*item.ProductID)
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "failed to check product")
			}
			if !isProductExist {
				return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Product ID %d not found", *item.ProductID))
			}
		}

		if item.GroupID != nil {
			isGroupExist, err := h.usecases.IsGroupIDExist(*item.GroupID)
			if err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "failed to check group")
			}
			if !isGroupExist {
				return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Group ID %d not found", *item.GroupID))
			}
		}

		items = append(items, domain.CartItem{
			CartID:    cart.ID,
			ProductID: item.ProductID,
			GroupID:   item.GroupID,
			Quantity:  item.Quantity,
		})
	}

	if err := h.usecases.UpdateItemInCartID(cart.ID, items); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to update cart item")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *CartItemHandler) Update(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id type")
	}
	cart, err := h.usecases.GetCartByUserID(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to get cart")
	}

	var req cartitemrequest.UpdateCartItemRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := req.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if req.ProductID != nil {
		isProductExist, err := h.usecases.IsProductIDExists(*req.ProductID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to check product")
		}
		if !isProductExist {
			return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Product ID %d not found", *req.ProductID))
		}
	}

	if req.GroupID != nil {
		isGroupExist, err := h.usecases.IsGroupIDExist(*req.GroupID)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "failed to check group")
		}
		if !isGroupExist {
			return fiber.NewError(fiber.StatusNotFound, fmt.Sprintf("Group ID %d not found", *req.GroupID))
		}
	}

	item := &domain.CartItem{
		CartID:    cart.ID,
		ProductID: req.ProductID,
		GroupID:   req.GroupID,
		Quantity:  req.Quantity,
	}

	if err := h.usecases.Update(item); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to update cart item")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *CartItemHandler) DeleteCartItem(c *fiber.Ctx) error {
	// 1. ดึง user_id จาก context
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user_id type")
	}

	// 2. ดึง cart ของ user
	cart, err := h.usecases.GetCartByUserID(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "failed to get cart")
	}

	// 3. ดึง cartItemID จาก params
	cartItemIDParam, err := strconv.Atoi(c.Params("cartItemID"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid cart item ID")
	}

	// 4. เรียก usecase/repo ลบ cart item
	if err := h.usecases.Delete(cart.ID, uint(cartItemIDParam)); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, "failed to delete cart item")
	}

	// 5. ส่ง status 200
	return c.SendStatus(fiber.StatusOK)
}
