package http

import (
	"github.com/gin-gonic/gin"
	"otel-prometheus-study/internal/logger"
)

type StoreProductController struct {
	// TODO : implement this
}

func (s *StoreProductController) Create(ctx *gin.Context) {
	// TODO : implement this
}

func (s *StoreProductController) List(ctx *gin.Context) {
	// TODO : implement this
}

func (s *StoreProductController) Get(ctx *gin.Context) {
	// TODO : implement this
}

func (s *StoreProductController) Update(ctx *gin.Context) {
	// TODO : implement this
}

func NewStoreProductController() StoreProductController {
	logger.LogInfo("initializing store product controller")
	return StoreProductController{}
}
