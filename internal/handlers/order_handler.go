package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vimudakorn/internal/domain"
	"github.com/vimudakorn/internal/usecases"
)

type OrderHandler struct {
	usecases *usecases.OrderUsecase
}

func NewOrderHandler(uc *usecases.OrderUsecase) *OrderHandler {
	return &OrderHandler{usecases: uc}
}

func (h *OrderHandler) Create(c *fiber.Ctx) error {
	var req domain.Order
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}
	if err := h.usecases.CreateOrder(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(req)
}

func (h *OrderHandler) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	o, err := h.usecases.GetOrderByID(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "order not found")
	}
	return c.JSON(o)
}

func (h *OrderHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	search := c.Query("search", "")
	sortBy := c.Query("sortBy", "id")
	orderBy := c.Query("orderBy", "asc")

	orders, total, err := h.usecases.GetOrders(page, limit, search, sortBy, orderBy)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(fiber.Map{
		"data":  orders,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

func (h *OrderHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var req domain.Order
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}
	req.ID = uint(id)
	if err := h.usecases.UpdateOrder(&req); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(req)
}

func (h *OrderHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.usecases.DeleteOrder(uint(id)); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
