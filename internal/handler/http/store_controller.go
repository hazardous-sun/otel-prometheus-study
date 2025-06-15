package http

import (
	"github.com/gin-gonic/gin"
	"otel-prometheus-study/internal/logger"
)

type StoreController struct {
	// TODO : implement this
}

func (s *StoreController) Create(ctx *gin.Context) {
	// TODO : implement this
}

func (s *StoreController) List(ctx *gin.Context) {
	// TODO : implement this
}

func (s *StoreController) Get(ctx *gin.Context) {
	// TODO : implement this
}

func (s *StoreController) Update(ctx *gin.Context) {
	// TODO : implement this
}

func NewStoreController() StoreController {
	logger.LogInfo("initializing store controller")
	return StoreController{}
}
