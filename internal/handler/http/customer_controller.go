package http

import "github.com/gin-gonic/gin"

type CustomerController struct {
	// TODO : implement this
}

func (c *CustomerController) Create(ctx *gin.Context) {
	// TODO : implement this
}

func (c *CustomerController) List(ctx *gin.Context) {
	// TODO : implement this
}

func (c *CustomerController) Get(ctx *gin.Context) {
	// TODO : implement this
}

func NewCustomerController() CustomerController {
	return CustomerController{}
}
