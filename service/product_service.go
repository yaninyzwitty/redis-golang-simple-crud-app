// internal/service/product_service.go
package service

import (
	"context"

	"github.com/yaninyzwitty/crud-sql/model"
	"github.com/yaninyzwitty/crud-sql/repository"
)

type ProductService interface {
	GetAllProducts(ctx context.Context) ([]model.Product, error)
	GetProduct(ctx context.Context, id string) (model.Product, error)
	CreateProduct(ctx context.Context, product model.Product) error
	UpdateProduct(ctx context.Context, id string, product model.Product) error
	DeleteProduct(ctx context.Context, id string) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	return s.repo.GetAllProducts(ctx)
}

func (s *productService) GetProduct(ctx context.Context, id string) (model.Product, error) {
	return s.repo.GetProduct(ctx, id)
}

func (s *productService) CreateProduct(ctx context.Context, product model.Product) error {
	return s.repo.CreateProduct(ctx, product)
}

func (s *productService) UpdateProduct(ctx context.Context, id string, product model.Product) error {
	return s.repo.UpdateProduct(ctx, id, product)
}

func (s *productService) DeleteProduct(ctx context.Context, id string) error {
	return s.repo.DeleteProduct(ctx, id)
}
