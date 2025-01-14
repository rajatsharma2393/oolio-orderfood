package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"orderfood/models"
	"testing"
)

var testProducts = []models.Product{
	{ID: "1", Name: "Cheese Pizza", Price: 10.99, Category: "Pizza"},
	{ID: "2", Name: "Veggie Burger", Price: 8.99, Category: "Burger"},
}

func newTestProductHandler() *ProductHandler {
	return &ProductHandler{products: testProducts}
}

func TestListProducts(t *testing.T) {
	handler := newTestProductHandler()

	w := httptest.NewRecorder()
	r := &http.Request{}

	handler.ListProducts(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("ListProducts handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var products []models.Product
	if err := json.NewDecoder(w.Body).Decode(&products); err != nil {
		t.Fatal(err)
	}

	if len(products) != len(testProducts) {
		t.Errorf("Expected %d products, got %d", len(testProducts), len(products))
	}
}

func TestGetProduct(t *testing.T) {
	handler := newTestProductHandler()

	tests := []struct {
		productID    string
		expectedCode int
		expectedName string
	}{
		{"1", http.StatusOK, "Cheese Pizza"},
		{"2", http.StatusOK, "Veggie Burger"},
		{"999", http.StatusNotFound, ""},
	}

	for _, tt := range tests {
		t.Run(tt.productID, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/product/"+tt.productID, nil)

			handler.GetProduct(w, req)

			if status := w.Code; status != tt.expectedCode {
				t.Errorf("GetProduct handler returned wrong status code: got %v want %v", status, tt.expectedCode)
			}

			if tt.expectedCode == http.StatusOK {
				var product models.Product
				if err := json.NewDecoder(w.Body).Decode(&product); err != nil {
					t.Fatal(err)
				}

				if product.Name != tt.expectedName {
					t.Errorf("GetProduct handler returned wrong product: got %v want %v", product.Name, tt.expectedName)
				}
			}
		})
	}
}
