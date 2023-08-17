package service

import (
	"errors"
	"product-rest/model"
	"product-rest/repository"
)

type ProductService interface{
	GetAllProducts(sortBy string, typeSort string) ([]model.Product, error)
	CreateNewProduct(model.Product) error
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *productService {
	return &productService{productRepository: productRepository}
}

func(s *productService) GetAllProducts(sortBy string, typeSort string) ([]model.Product, error) {

	products, err := s.productRepository.GetAllProducts(sortBy, typeSort)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func(s *productService) CreateNewProduct(product model.Product) error {
	
	if err := isValidProduct(product); err != nil {
		return err
	}

	err := s.productRepository.InsertProduct(product)
	if err != nil {
		return err
	}

	return nil
}

func isValidProduct(product model.Product) error {
	if product.Name == "" {
		return errors.New("product name cannot be empty")
	}

	if product.Price < 1 {
		return errors.New("price cannot be less than 1")
	}

	if product.Quantity < 1 {
		return errors.New("quantity cannot be less than 1")
	}

	return nil
}