// internal/repository/product_repository.go
package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/yaninyzwitty/crud-sql/model"
)

type ProductRepository interface {
	GetAllProducts(ctx context.Context) ([]model.Product, error)
	GetProduct(ctx context.Context, id string) (model.Product, error)
	CreateProduct(ctx context.Context, product model.Product) error
	UpdateProduct(ctx context.Context, id string, product model.Product) error
	DeleteProduct(ctx context.Context, id string) error
}

type productRepository struct {
	db *redis.Client
}

var ErrorNotExists = errors.New("product does not exist")

func NewProductRepository(ctx context.Context, db *redis.Client) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	// Redis logic to get all products

	res := r.db.SScan(ctx, "products", 0, "", 0)
	keys, _, err := res.Result()
	if err != nil {
		return []model.Product{}, fmt.Errorf("failed to get products keys: %w", err)
	}

	if len(keys) == 0 {
		return []model.Product{}, nil
	}

	xs, err := r.db.MGet(ctx, keys...).Result()

	if err != nil {
		return []model.Product{}, fmt.Errorf("failed to get products: %w", err)
	}
	products := make([]model.Product, len(xs))

	for i, x := range xs {
		x := x.(string)
		var product model.Product
		err := json.Unmarshal([]byte(x), &product)
		if err != nil {
			return []model.Product{}, fmt.Errorf("failed to decode product: %w", err)
		}

		products[i] = product

	}

	return products, nil
}

func (r *productRepository) GetProduct(ctx context.Context, id string) (model.Product, error) {
	// Redis logic to get a product by id

	key := fmt.Sprintf("product:%s", id)
	val, err := r.db.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return model.Product{}, ErrorNotExists

	} else if err != nil {
		return model.Product{}, fmt.Errorf("failed to get product: %w", err)
	}

	// here you convert val (json data) to of model type

	var product model.Product
	err = json.Unmarshal([]byte(val), &product)
	if err != nil {
		return model.Product{}, fmt.Errorf("failed to unmarshal product: %w", err)
	}

	return product, nil

}

func (r *productRepository) CreateProduct(ctx context.Context, product model.Product) error {
	// Redis logic to create a product
	data, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("error marshalling product: %v", err)
	}
	key := fmt.Sprintf("product:%s", product.ID)

	res := r.db.SetNX(ctx, key, data, 0)
	if err := res.Err(); err != nil {
		return fmt.Errorf("failed to set product: %w", err)
	}

	return nil

}

func (r *productRepository) UpdateProduct(ctx context.Context, id string, product model.Product) error {
	// Redis logic to update a product
	data, err := json.Marshal(product)
	if err != nil {
		return fmt.Errorf("error marshalling product: %v", err)
	}
	key := fmt.Sprintf("product:%s", id)
	err = r.db.SetXX(ctx, key, data, 0).Err()
	if errors.Is(err, redis.Nil) {
		return ErrorNotExists
	} else if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}

	return nil
}

func (r *productRepository) DeleteProduct(ctx context.Context, id string) error {
	// Redis logic to delete a product
	key := fmt.Sprintf("product:%s", id)

	err := r.db.Del(ctx, key).Err()

	if errors.Is(redis.Nil, err) {
		return ErrorNotExists
	} else if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	return nil
}
