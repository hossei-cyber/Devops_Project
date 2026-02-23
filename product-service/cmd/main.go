package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/robinlieb/devops-lecture-project-2026/product-service/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/products", handlers.ProductListHandler)
	mux.HandleFunc("/products/{id}", handlers.ProductDetailHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("product-service running on :%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}