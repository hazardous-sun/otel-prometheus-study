package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"otel-prometheus-study/internal/app/product"
	"otel-prometheus-study/internal/logger"
)

type ProductController struct {
}

func NewProductController() ProductController {
	logger.LogInfo("initializing product controller")
	return ProductController{}
}

func (pc *ProductController) Create(c *gin.Context) {
	var json struct {
		Name  string `json:"name"`
		Price string `json:"price"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newProduct, err := product.CreateProduct(json.Name, json.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}

func (pc *ProductController) List(c *gin.Context) {
	products, err := product.ListProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
