package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/robinlieb/devops-lecture-project-2026/checkout-service/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/checkout/placeorder", handlers.CheckoutPlaceOrderHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("checkout-service running on :%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}