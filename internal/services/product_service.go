package services

import (
    // "errors"

    "github.com/Syed-Rehan-21/GO-lang-/internal/models"
    "github.com/Syed-Rehan-21/GO-lang-/internal/repository"
)

// ProductService handles business logic for products
type ProductService struct {
    Repo *repository.ProductRepository
}

// NewProductService creates a new instance of ProductService
func NewProductService(repo *repository.ProductRepository) *ProductService {
    return &ProductService{Repo: repo}
}

// GetAllProducts fetches all products
func (s *ProductService) GetAllProducts() ([]models.Product, error) {
    return s.Repo.GetAllProducts()
}

// GetProductByID fetches a single product by ID
func (s *ProductService) GetProductByID(id int64) (*models.Product, error) {
    return s.Repo.GetProductByID(id)
}