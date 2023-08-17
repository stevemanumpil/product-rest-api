package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"product-rest/model"

	"github.com/gin-gonic/gin"
)

func(h *handler) GetProduct(c *gin.Context) {

	var sortBy, typeSort string
	
	if c.Query("sort_by") != "" {
		sortBy = c.Query("sort_by")
	}

	if c.Query("sort_type") != "" {
		typeSort = c.Query("sort_type")
	}

	products, err := h.productService.GetAllProducts(sortBy, typeSort)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

func(h *handler) PostProduct(c *gin.Context) {

	byteBody, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	requestBody := bytes.NewBuffer(byteBody)
	if len(requestBody.String()) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "body cannot empty",
		})
		return
	}
	
	var product model.Product
	err = json.Unmarshal(byteBody, &product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	err = h.productService.CreateNewProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product inserted",
	})
}