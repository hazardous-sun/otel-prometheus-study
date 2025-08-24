package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"otel-prometheus-study/internal/app/store_product"
	"otel-prometheus-study/internal/logger"
)

type StoreProductController struct {
}

func NewStoreProductController() StoreProductController {
	logger.LogInfo("initializing store product controller")
	return StoreProductController{}
}

func (spc *StoreProductController) Create(c *gin.Context) {
	var json struct {
		StoreID   int    `json:"store_id"`
		ProductID int    `json:"product_id"`
		Price     string `json:"price"`
		Quantity  int    `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newStoreProduct, err := store_product.CreateStoreProduct(json.StoreID, json.ProductID, json.Price, json.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newStoreProduct)
}

func (spc *StoreProductController) List(c *gin.Context) {
	storeProducts, err := store_product.ListStoreProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, storeProducts)
}
