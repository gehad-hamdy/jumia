package routes

import (
	"Flash/controllers"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func CustomerRoutes(route *gin.Engine) {
	customerController := controllers.GetCustomerController()
	router := route.Group("customer/")
	{
		router.GET("/", customerController.ListCustomerData)
	}
}
