package repository

import (
    "database/sql"
    "errors"

    "github.com/Syed-Rehan-21/GO-lang-/internal/models"
)

// ProductRepository handles database operations for products
type ProductRepository struct {
    DB *sql.DB
}

// NewProductRepository creates a new instance of ProductRepository
func NewProductRepository(db *sql.DB) *ProductRepository {
    return &ProductRepository{DB: db}
}

// GetAllProducts retrieves all products from the database
func (r *ProductRepository) GetAllProducts() ([]models.Product, error) {
    query := "SELECT id, name, quantity, price FROM products"
    rows, err := r.DB.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []models.Product
    for rows.Next() {
        var product models.Product
        if err := rows.Scan(&product.ID, &product.Name, &product.Quantity, &product.Price); err != nil {
            return nil, err
        }
        products = append(products, product)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return products, nil
}

// GetProductByID retrieves a single product by ID
func (r *ProductRepository) GetProductByID(id int64) (*models.Product, error) {
    query := "SELECT id, name, quantity, price FROM products WHERE id = $1"
    row := r.DB.QueryRow(query, id)

    var product models.Product
    if err := row.Scan(&product.ID, &product.Name, &product.Quantity, &product.Price); err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.New("product not found")
        }
        return nil, err
    }

    return &product, nil
}