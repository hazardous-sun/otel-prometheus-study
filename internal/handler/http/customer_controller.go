package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"otel-prometheus-study/internal/app/customer"
	"otel-prometheus-study/internal/logger"
)

type CustomerController struct {
}

func NewCustomerController() CustomerController {
	logger.LogInfo("initializing customer controller")
	return CustomerController{}
}

func (cc *CustomerController) Create(c *gin.Context) {
	var json struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newCustomer, err := customer.CreateCustomer(json.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newCustomer)
}

func (cc *CustomerController) List(c *gin.Context) {
	customers, err := customer.ListCustomers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customers)
}
