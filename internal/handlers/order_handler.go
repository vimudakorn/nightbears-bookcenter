package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/vimudakorn/internal/domain"
	orderrequest "github.com/vimudakorn/internal/request/order_request"
	"github.com/vimudakorn/internal/usecases"
)

type OrderHandler struct {
	usecases *usecases.OrderUsecase
}

func NewOrderHandler(uc *usecases.OrderUsecase) *OrderHandler {
	return &OrderHandler{usecases: uc}
}

func (h *OrderHandler) GetByUserID(c *fiber.Ctx) error {
	userID, err := strconv.Atoi(c.Params("user_id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid user ID")
	}

	cart, err := h.usecases.GetOryderByUserID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to find cart by user ID"})
	}
	return c.JSON(fiber.Map{
		"data": cart,
	})
}

//	func (h *OrderHandler) Create(c *fiber.Ctx) error {
//		var req domain.Order
//		if err := c.BodyParser(&req); err != nil {
//			return fiber.NewError(fiber.StatusBadRequest, "invalid body")
//		}
//		if err := h.usecases.CreateOrder(&req); err != nil {
//			return fiber.NewError(fiber.StatusBadRequest, err.Error())
//		}
//		return c.Status(fiber.StatusCreated).JSON(req)
//	}
func (h *OrderHandler) Create(c *fiber.Ctx) error {
	// ดึง user_id จาก context (เช่น JWT)
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return fiber.NewError(fiber.StatusUnauthorized, "invalid user_id")
	}

	var req orderrequest.CreateOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}

	// สร้าง order จากข้อมูล request และ user_id จาก context
	order := domain.Order{
		UserID:     userID,
		TotalPrice: req.TotalPrice,
		Status:     req.Status,
	}

	// เพิ่ม order items
	for _, item := range req.Items {
		order.Items = append(order.Items, domain.OrderItem{
			ProductID:       item.ProductID,
			GroupID:         item.GroupID,
			Quantity:        item.Quantity,
			PriceAtPurchase: item.Price,
		})
	}

	// เรียก usecase สร้าง order
	if err := h.usecases.CreateOrder(&order); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(order)
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

// func (h *OrderHandler) Update(c *fiber.Ctx) error {
// 	id, _ := strconv.Atoi(c.Params("id"))
// 	var req domain.Order
// 	if err := c.BodyParser(&req); err != nil {
// 		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
// 	}
// 	req.ID = uint(id)
// 	if err := h.usecases.UpdateOrder(&req); err != nil {
// 		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
// 	}
// 	return c.JSON(req)
// }

func (h *OrderHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var req orderrequest.UpdateOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid body")
	}

	// อัปเดตเฉพาะ field ของ order
	updateData := map[string]interface{}{}
	if req.Status != "" {
		updateData["status"] = req.Status
	}
	if len(updateData) > 0 {
		if err := h.usecases.UpdateOrderFields(uint(id), updateData); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	// อัปเดต order items ถ้ามี
	if len(req.Items) > 0 {
		var items []domain.OrderItem
		for _, item := range req.Items {
			items = append(items, domain.OrderItem{
				ProductID:       item.ProductID,
				GroupID:         item.GroupID,
				Quantity:        item.Quantity,
				PriceAtPurchase: item.Price,
				OrderID:         uint(id),
			})
		}
		if err := h.usecases.UpdateItemsInOrderID(uint(id), items); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "order updated successfully"})
}

func (h *OrderHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.usecases.DeleteOrder(uint(id)); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
