package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"otel-prometheus-study/internal/logger"
)

func NewRouter(
	customerCtrl CustomerController,
	productCtrl ProductController,
	stockCtrl StockController,
	storeCtrl StoreController,
	storeProductCtrl StoreProductController,
) *gin.Engine {
	logger.LogInfo("initializing router")
	router := gin.New()

	// Global middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// TODO : Implement custom middleware for telemetry
	// TODO : Implement custom middleware for tracing
	// TODO : Implement custom middleware for logging
	// TODO : Implement custom timeout

	// System endpoints
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"}) // TODO : Implement custom health check
	})

	// Product routes
	productGroup := router.Group("/products")
	{
		productGroup.POST("", productCtrl.Create)
		productGroup.GET("", productCtrl.List)
		productGroup.GET("/:id", productCtrl.Get)
		productGroup.PUT("/:id", productCtrl.Update)
	}

	// Customer routes
	customerGroup := router.Group("/customers")
	{
		customerGroup.POST("", customerCtrl.Create)
		customerGroup.GET("", customerCtrl.List)
		customerGroup.GET("/:id", customerCtrl.Get)
	}

	// Stock routes
	stockGroup := router.Group("/stocks")
	{
		stockGroup.POST("", stockCtrl.Create)
		stockGroup.GET("", stockCtrl.List)
		stockGroup.GET("/:id", stockCtrl.Get)
		stockGroup.PUT("/:id", stockCtrl.Update)
	}

	// Store routes
	storeGroup := router.Group("/stores")
	{
		storeGroup.POST("", storeCtrl.Create)
		storeGroup.GET("", storeCtrl.List)
		storeGroup.GET("/:id", storeCtrl.Get)
		storeGroup.PUT("/:id", storeCtrl.Update)
	}

	// Store Product routes
	storeProductGroup := router.Group("/store_products")
	{
		storeProductGroup.POST("", storeProductCtrl.Create)
		storeProductGroup.GET("", storeProductCtrl.List)
		storeProductGroup.GET("/:id", storeProductCtrl.Get)
		storeProductGroup.PUT("/:id", storeProductCtrl.Update)
	}

	return router
}
