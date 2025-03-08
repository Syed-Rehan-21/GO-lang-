// @title Product API
// @version 1.0
// @description API for managing products
// @host localhost:8080
// @BasePath /api
package main

import (
    "fmt"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "github.com/sirupsen/logrus"

    "github.com/Syed-Rehan-21/GO-lang-/configs"
    "github.com/Syed-Rehan-21/GO-lang-/internal/handlers"
    "github.com/Syed-Rehan-21/GO-lang-/internal/repository"
    "github.com/Syed-Rehan-21/GO-lang-/internal/services"
    "github.com/Syed-Rehan-21/GO-lang-/pkg/utils"
)

var logger = logrus.New()

func init() {
    if err := godotenv.Load(); err != nil {
        logger.Warn("No .env file found, using environment variables directly.")
    }
    logger.SetFormatter(&logrus.JSONFormatter{})
    logger.SetOutput(os.Stdout)
    logger.SetLevel(logrus.InfoLevel)
}

func main() {
    config, err := configs.LoadConfig()
    if err != nil {
        logger.Fatalf("Failed to load configuration: %v", err)
    }

    db, err := utils.InitializeDB(config.Database)
    if err != nil {
        logger.Fatalf("Failed to initialize database: %v", err)
    }
    defer db.Close()

    if err := utils.CreateProductsTable(db); err != nil {
        logger.Fatalf("Failed to create products table: %v", err)
    }

    router := mux.NewRouter()

    // Initialize layers
    repo := repository.NewProductRepository(db)
    service := services.NewProductService(repo)
    handler := handlers.NewProductHandler(service, logger)

    // Define routes
    router.HandleFunc("/", 
    // @Summary Root endpoint
    // @Description Returns a simple welcome message to verify the server is running.
    // @Tags root
    // @Produce json
    // @Success 200 {object} map[string]string "Welcome message"
    // @Router / [get]
    func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        fmt.Fprintln(w, `{"message": "Hello Go"}`)
    }).Methods("GET")

    // Create API subrouter
    apiRouter := router.PathPrefix("/api").Subrouter()

    // Define API routes
    apiRouter.HandleFunc("/products", handler.GetAllProductsHandler).Methods("GET")
    apiRouter.HandleFunc("/products/{id}", handler.GetProductByIDHandler).Methods("GET")
    apiRouter.HandleFunc("/products", handler.CreateProductHandler).Methods("POST")
    apiRouter.HandleFunc("/products/{id}", handler.UpdateProductHandler).Methods("PUT")
    apiRouter.HandleFunc("/products/{id}", handler.DeleteProductHandler).Methods("DELETE")

    // Serve Swagger UI and OpenAPI documentation
    fs := http.FileServer(http.Dir("./api/docs"))
    router.PathPrefix("/api/docs/").Handler(http.StripPrefix("/api/docs/", fs))

    // Start the HTTP server
    addr := fmt.Sprintf(":%s", config.APIPort)
    logger.Infof("Starting server on port %s", config.APIPort)
    if err := http.ListenAndServe(addr, router); err != nil {
        logger.Fatalf("Server failed to start: %v", err)
    }
}