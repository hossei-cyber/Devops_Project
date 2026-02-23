package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/robinlieb/devops-lecture-project-2026/auth-service/internal/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/auth/login", handlers.AuthLoginHandler)
	mux.HandleFunc("/auth/logout", handlers.AuthLogoutHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("auth-service running on :%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}