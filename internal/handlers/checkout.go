package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"github.com/robinlieb/devops-lecture-project-2026/internal/models"
	"github.com/robinlieb/devops-lecture-project-2026/pkg/auth"
)
func CheckoutPlaceOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"Missing Authorization header"}`))
		return
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"Authorization header must use Bearer scheme"}`))
		return
	}

	tokenString := strings.TrimPrefix(authHeader, bearerPrefix)

	if !auth.VerifyToken(tokenString) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"Invalid token"}`))
		return
	}

	orderID := "ORD-" + fmt.Sprintf("%d", len(models.Products)*100+42)

	response := map[string]interface{}{
		"message":  "Order placed successfully",
		"order_id": orderID,
		"status":   "confirmed",
	}

	jsonResponse, _ := json.Marshal(response)
	w.Write(jsonResponse)
}