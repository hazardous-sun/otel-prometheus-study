package http

import (
	"github.com/gin-gonic/gin"
	"otel-prometheus-study/internal/logger"
)

type StockController struct {
	// TODO : implement this
}

func (s *StockController) Create(ctx *gin.Context) {
	// TODO : implement this
}

func (s *StockController) List(ctx *gin.Context) {
	// TODO : implement this
}

func (s *StockController) Get(ctx *gin.Context) {
	// TODO : implement this
}

func (s *StockController) Update(ctx *gin.Context) {
	// TODO : implement this
}

func NewStockController() StockController {
	logger.LogInfo("initializing stock controller")
	return StockController{}
}
