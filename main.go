package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"orderfood/handlers"
	"orderfood/middleware"
	"orderfood/services"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	couponService := services.NewCouponService()

	orderHandler := handlers.NewOrderHandler(couponService)
	productHandler := handlers.NewProductHandler()

	http.HandleFunc("/product", productHandler.ListProducts)
	http.HandleFunc("/product/", productHandler.GetProduct)
	http.HandleFunc("/order", orderHandler.PlaceOrder)

	securedHandler := middleware.AuthMiddleware(http.DefaultServeMux)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", securedHandler))
}
