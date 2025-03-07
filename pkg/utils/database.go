package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver

	"github.com/Syed-Rehan-21/GO-lang-/configs" // Update with your actual module name
)

// InitializeDB establishes a connection to the PostgreSQL database
func InitializeDB(config configs.DBConfig) (*sql.DB, error) {
	// Create the connection string
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName,
	)

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not connect to database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping database: %w", err)
	}

	return db, nil
}

// CreateProductsTable creates the products table if it doesn't exist
func CreateProductsTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		quantity INT NOT NULL,
		price FLOAT NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("could not create products table: %w", err)
	}

	return nil
}