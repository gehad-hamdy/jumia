package controllers

import (
	"Flash/resources"
	"Flash/services"
	"github.com/gin-gonic/gin"
)

type CustomersController struct {
	customerService services.CustomerService
}

func GetCustomerController() *CustomersController {
	return &CustomersController{customerService: services.GetCustomerService()}
}

func (controller CustomersController) ListCustomerData(context *gin.Context) {
	listNumbers := controller.customerService.ListPhoneNumbers()
	data := controller.customerService.CheckValidationNumber(listNumbers)
	response := resources.GetSuccess200Resource(data, "Success")
	context.JSON(response.GetStatus(), response.GetData())
}