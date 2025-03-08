package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/sirupsen/logrus"

    "github.com/Syed-Rehan-21/GO-lang-/internal/models"
    "github.com/Syed-Rehan-21/GO-lang-/internal/services"
)

// ProductHandler handles HTTP requests for products
type ProductHandler struct {
    Service *services.ProductService
    Logger  *logrus.Logger
}

// NewProductHandler creates a new instance of ProductHandler
func NewProductHandler(service *services.ProductService, logger *logrus.Logger) *ProductHandler {
    return &ProductHandler{Service: service, Logger: logger}
}

// @Summary Get all products
// @Description Retrieves a list of all products in the inventory
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Failure 500 {object} map[string]string
// @Router /products [get]
// GetAllProductsHandler handles GET /products
func (h *ProductHandler) GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
    products, err := h.Service.GetAllProducts()
    if err != nil {
        h.Logger.Errorf("Failed to fetch products: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(products)
}

// @Summary Get a product by ID
// @Description Retrieves a single product by its unique ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Failure 404 {object} map[string]string "Product not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /products/{id} [get]
// GetProductByIDHandler handles GET /products/{id}
func (h *ProductHandler) GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.ParseInt(vars["id"], 10, 64)
    if err != nil {
        h.Logger.Errorf("Invalid product ID: %v", err)
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    product, err := h.Service.GetProductByID(id)
    if err != nil {
        if err.Error() == "product not found" {
            h.Logger.Warnf("Product with ID %d not found", id)
            http.Error(w, "Product not found", http.StatusNotFound)
            return
        }
        h.Logger.Errorf("Failed to fetch product: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(product)
}

// @Summary Create a new product
// @Description Adds a new product to the inventory
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.ProductInput true "Product details"
// @Success 201 {object} models.Product
// @Failure 400 {object} map[string]string
// @Router /products [post]
// CreateProductHandler handles POST /products
func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
    // Parse the request body into a ProductInput struct
    var input models.ProductInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        h.Logger.Errorf("Invalid request body: %v", err)
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Call the service to create the product
    product, err := h.Service.CreateProduct(input)
    if err != nil {
        h.Logger.Errorf("Failed to create product: %v", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Return the created product as JSON
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated) // 201 Created
    json.NewEncoder(w).Encode(product)
}

// @Summary Update a product
// @Description Updates an existing product in the inventory
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.ProductInput true "Updated product details"
// @Success 200 {object} models.Product
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 404 {object} map[string]string "Product not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /products/{id} [put]
// UpdateProductHandler handles PUT /products/{id}
func (h *ProductHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
    // Extract the product ID from the URL path
    vars := mux.Vars(r)
    id, err := strconv.ParseInt(vars["id"], 10, 64)
    if err != nil {
        h.Logger.Errorf("Invalid product ID: %v", err)
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    // Parse the request body into a ProductInput struct
    var input models.ProductInput
    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        h.Logger.Errorf("Invalid request body: %v", err)
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Call the service to update the product
    product, err := h.Service.UpdateProduct(id, input)
    if err != nil {
        if err.Error() == "product not found" {
            h.Logger.Warnf("Product with ID %d not found", id)
            http.Error(w, "Product not found", http.StatusNotFound)
            return
        }
        h.Logger.Errorf("Failed to update product: %v", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Return the updated product as JSON
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK) // 200 OK
    json.NewEncoder(w).Encode(product)
}

// @Summary Delete a product
// @Description Deletes a product from the inventory
// @Tags products
// @Param id path int true "Product ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string "Product not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /products/{id} [delete]
// DeleteProductHandler handles DELETE /products/{id}
func (h *ProductHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
    // Extract the product ID from the URL path
    vars := mux.Vars(r)
    id, err := strconv.ParseInt(vars["id"], 10, 64)
    if err != nil {
        h.Logger.Errorf("Invalid product ID: %v", err)
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    // Call the service to delete the product
    err = h.Service.DeleteProduct(id)
    if err != nil {
        if err.Error() == "product not found" {
            h.Logger.Warnf("Product with ID %d not found", id)
            http.Error(w, "Product not found", http.StatusNotFound)
            return
        }
        h.Logger.Errorf("Failed to delete product: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Return a success response
    w.WriteHeader(http.StatusNoContent) // 204 No Content
}