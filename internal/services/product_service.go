package services

import (
    "errors"

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

// CreateProduct validates and creates a new product
func (s *ProductService) CreateProduct(input models.ProductInput) (*models.Product, error) {
    // Validate input
    if input.Quantity <= 0 {
        return nil, errors.New("quantity must be greater than zero")
    }
    if input.Price <= 0 {
        return nil, errors.New("price must be greater than zero")
    }

    // Call repository to create the product
    return s.Repo.CreateProduct(input)
}

// UpdateProduct validates and updates an existing product
func (s *ProductService) UpdateProduct(id int64, input models.ProductInput) (*models.Product, error) {
    // Validate input
    if input.Quantity <= 0 {
        return nil, errors.New("quantity must be greater than zero")
    }
    if input.Price <= 0 {
        return nil, errors.New("price must be greater than zero")
    }

    // Call repository to update the product
    return s.Repo.UpdateProduct(id, input)
}

// DeleteProduct deletes a product by ID
func (s *ProductService) DeleteProduct(id int64) error {
    // Call repository to delete the product
    return s.Repo.DeleteProduct(id)
}