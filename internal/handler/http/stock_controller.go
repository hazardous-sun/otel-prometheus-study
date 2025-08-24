package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"otel-prometheus-study/internal/app/stock"
	"otel-prometheus-study/internal/logger"
)

type StockController struct {
}

func NewStockController() StockController {
	logger.LogInfo("initializing stock controller")
	return StockController{}
}

func (sc *StockController) Create(c *gin.Context) {
	var json struct {
		ProductID int `json:"product_id"`
		Quantity  int `json:"quantity"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newStock, err := stock.CreateStock(json.ProductID, json.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newStock)
}

func (sc *StockController) List(c *gin.Context) {
	stocks, err := stock.ListStocks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stocks)
}
