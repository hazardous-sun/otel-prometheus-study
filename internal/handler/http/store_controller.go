package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"otel-prometheus-study/internal/app/store"
	"otel-prometheus-study/internal/logger"
)

type StoreController struct {
}

func NewStoreController() StoreController {
	logger.LogInfo("initializing store controller")
	return StoreController{}
}

func (sc *StoreController) Create(c *gin.Context) {
	var json struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newStore, err := store.CreateStore(json.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newStore)
}

func (sc *StoreController) List(c *gin.Context) {
	stores, err := store.ListStores()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stores)
}
