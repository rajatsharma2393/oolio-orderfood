package models

import "github.com/shopspring/decimal"

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

type OrderReq struct {
	CouponCode string `json:"couponCode"`
	Items      []struct {
		ProductID string `json:"productId"`
		Quantity  int    `json:"quantity"`
	} `json:"items"`
}

type Order struct {
	ID    string      `json:"id"`
	Items []OrderItem `json:"items"`
}

type OrderItem struct {
	ProductID string          `json:"productId"`
	Quantity  int             `json:"quantity"`
	Price     decimal.Decimal `json:"price"`
	Name      string          `json:"name"`
	Category  string          `json:"category"`
}
