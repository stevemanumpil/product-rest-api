package handler

import (
	"product-rest/service"
)

type handler struct {
	productService service.ProductService
}

func NewHandler(services *service.Services) *handler {
	return &handler{productService: services.ProductService}
}