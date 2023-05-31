// domain/product.go
package domain

// Product represents a product in WooCommerce
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// ProductService defines the interface for managing products
type ProductService interface {
	GetProducts() ([]Product, error)
	GetProductByID(id int) (*Product, error)
	CreateProduct(product Product) (*Product, error)
	UpdateProduct(id int, product Product) (*Product, error)
	DeleteProduct(id int) error
}
