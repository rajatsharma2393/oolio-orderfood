package handlers

import (
	"encoding/json"
	"net/http"
	"orderfood/models"
	"strings"
)

type ProductHandler struct {
	products []models.Product
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{products: products}
}

var products = []models.Product{
	{ID: "1", Name: "Cheese Pizza", Price: 10.99, Category: "Pizza"},
	{ID: "2", Name: "Veggie Burger", Price: 8.99, Category: "Burger"},
}

func (h *ProductHandler) ListProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.products)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	productID := segments[2]
	for _, product := range h.products {
		if product.ID == productID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(product)
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}
