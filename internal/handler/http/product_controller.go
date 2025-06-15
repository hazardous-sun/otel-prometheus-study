package http

import (
	"github.com/gin-gonic/gin"
	"otel-prometheus-study/internal/logger"
)

type ProductController struct {
	// TODO : implement this
}

func (p *ProductController) Create(ctx *gin.Context) {
	// TODO : implement this
}

func (p *ProductController) List(ctx *gin.Context) {
	// TODO : implement this
}

func (p *ProductController) Get(ctx *gin.Context) {
	// TODO : implement this
}

func (p *ProductController) Update(ctx *gin.Context) {
	// TODO : implement this
}

func NewProductController() ProductController {
	logger.LogInfo("initializing product controller")
	return ProductController{}
}
