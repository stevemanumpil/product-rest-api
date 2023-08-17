package app

import (
	"fmt"
	"product-rest/config"
	"product-rest/repository"
	"product-rest/router"
	"product-rest/service"
)

func Execute() {

	cnf := config.GetConfig()
	db := config.DBProvider(cnf.DB)

	productRepository := repository.NewProductRepository(db)

	productService := service.NewProductService(productRepository)

	g := router.NewRouter(&service.Services{
		ProductService: productService,
	})

	addr := fmt.Sprintf("%s:%v", cnf.App.Host, cnf.App.Port)

	g.Run(addr)
}