package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"net/http"
	"orderfood/models"
	"orderfood/services"
)

type OrderHandler struct {
	couponService *services.CouponService
}

func NewOrderHandler(couponService *services.CouponService) *OrderHandler {
	return &OrderHandler{couponService: couponService}
}

func (h *OrderHandler) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	var orderReq models.OrderReq
	if err := json.NewDecoder(r.Body).Decode(&orderReq); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if orderReq.CouponCode != "" {
		if !h.couponService.IsValidCoupon(orderReq.CouponCode) {
			http.Error(w, "Invalid coupon code", http.StatusBadRequest)
			return
		}
	}

	orderItems := make([]models.OrderItem, 0, len(orderReq.Items))

	for _, reqItem := range orderReq.Items {
		var product *models.Product
		for _, p := range products {
			if p.ID == reqItem.ProductID {
				product = &p
				break
			}
		}

		if product == nil {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}

		orderItems = append(orderItems, models.OrderItem{
			ProductID: reqItem.ProductID,
			Quantity:  reqItem.Quantity,
			Price:     decimal.NewFromFloat(product.Price * float64(reqItem.Quantity)),
			Name:      product.Name,
			Category:  product.Category,
		})
	}

	order := models.Order{
		ID:    uuid.New().String(),
		Items: orderItems,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}
