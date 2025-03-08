package repository

import (
    "database/sql"
    "errors"
	"fmt"

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

// CreateProduct inserts a new product into the database
func (r *ProductRepository) CreateProduct(product models.ProductInput) (*models.Product, error) {
    query := `
        INSERT INTO products (name, quantity, price)
        VALUES ($1, $2, $3)
        RETURNING id, name, quantity, price
    `
    var newProduct models.Product
    err := r.DB.QueryRow(query, product.Name, product.Quantity, product.Price).
        Scan(&newProduct.ID, &newProduct.Name, &newProduct.Quantity, &newProduct.Price)
    if err != nil {
        return nil, fmt.Errorf("failed to create product: %w", err)
    }
    return &newProduct, nil
}

// UpdateProduct updates an existing product in the database
func (r *ProductRepository) UpdateProduct(id int64, product models.ProductInput) (*models.Product, error) {
    query := `
        UPDATE products
        SET name = $1, quantity = $2, price = $3
        WHERE id = $4
        RETURNING id, name, quantity, price
    `
    var updatedProduct models.Product
    err := r.DB.QueryRow(query, product.Name, product.Quantity, product.Price, id).
        Scan(&updatedProduct.ID, &updatedProduct.Name, &updatedProduct.Quantity, &updatedProduct.Price)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errors.New("product not found")
        }
        return nil, fmt.Errorf("failed to update product: %w", err)
    }
    return &updatedProduct, nil
}

// DeleteProduct deletes a product from the database by ID
func (r *ProductRepository) DeleteProduct(id int64) error {
    query := "DELETE FROM products WHERE id = $1"
    result, err := r.DB.Exec(query, id)
    if err != nil {
        return fmt.Errorf("failed to delete product: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to check rows affected: %w", err)
    }

    if rowsAffected == 0 {
        return errors.New("product not found")
    }

    return nil
}