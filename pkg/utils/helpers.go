package utils

import "github.com/robinlieb/devops-lecture-project-2026/internal/models"

func FindProductByID(products []models.Product, id int) *models.Product {
	for _, product := range products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}
