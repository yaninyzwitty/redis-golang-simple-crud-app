// internal/model/product.go
package model

type Product struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	StockQuantity int     `json:"stock_quantity"`
}

type Category struct {
	ID   string `json:"category_id"`
	Name string `json:"name"`
}
