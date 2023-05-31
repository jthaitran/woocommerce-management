// service/product_service.go
package service

import (
	"github.com/jthaitran/woocommerce-management/domain"
)

// ProductService is responsible for managing products
type ProductService struct {
	repo domain.ProductRepository
}

// NewProductService creates a new instance of ProductService
func NewProductService(repo domain.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

// GetProducts retrieves all products
func (s *ProductService) GetProducts() ([]domain.Product, error) {
	return s.repo.GetProducts()
}

// GetProductByID retrieves a product by ID
func (s *ProductService) GetProductByID(id int) (*domain.Product, error) {
	return s.repo.GetProductByID(id)
}

// CreateProduct creates a new product
func (s *ProductService) CreateProduct(product domain.Product) (*domain.Product, error) {
	return s.repo.CreateProduct(product)
}

// UpdateProduct updates an existing product
func (s *ProductService) UpdateProduct(id int, product domain.Product) (*domain.Product, error) {
	return s.repo.UpdateProduct(id, product)
}

// DeleteProduct deletes a product by ID
func (s *ProductService) DeleteProduct(id int) error {
	return s.repo.DeleteProduct(id)
}
