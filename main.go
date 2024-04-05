// cmd/your-microservice/main.go
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/redis/go-redis/v9"
	"github.com/yaninyzwitty/crud-sql/repository"
	"github.com/yaninyzwitty/crud-sql/service"
	"github.com/yaninyzwitty/crud-sql/transport"
)

func main() {
	// Initialize Redis client
	ctx := context.Background()
	db := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Initialize dependencies
	productRepo := repository.NewProductRepository(ctx, db) // Pass context here
	productService := service.NewProductService(productRepo)
	productHandler := transport.NewProductHandler(productService)

	// Setup router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/products", productHandler.GetAllProducts)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Post("/products", productHandler.CreateProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	// Start server
	log.Println("Starting server on :3000...")
	log.Fatal(http.ListenAndServe(":3000", r))
}
