package router

import (
	"product-rest/handler"
	"product-rest/service"

	"github.com/gin-gonic/gin"
)

func NewRouter(s *service.Services) *gin.Engine {

	route := gin.Default()
	route.HandleMethodNotAllowed = true

	h := handler.NewHandler(s)

	route.GET("/product", h.GetProduct)
	route.POST("/product", h.PostProduct)

	return route

}