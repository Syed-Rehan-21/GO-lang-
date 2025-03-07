package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
    "github.com/sirupsen/logrus"

    // "github.com/Syed-Rehan-21/GO-lang-/internal/models"
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