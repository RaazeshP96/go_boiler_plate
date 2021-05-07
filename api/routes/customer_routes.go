package routes

import (
	"log"

	customer_controller "prototype2/api/controller/customer"
	customer_repository "prototype2/api/repository/customer"
	customer_service "prototype2/api/service/customer"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CustomerRoutes(route *gin.RouterGroup, db *gorm.DB) {
	customerRepository := customer_repository.NewCustomerRepository(db)
	if err := customerRepository.Migrate(); err != nil {
		log.Fatal("Customer migrate err", err)
	}
	customerService := customer_service.NewCustomerService(customerRepository)
	customerController := customer_controller.NewCustomerController(customerService)

	route.GET("/", customerController.GetCustomers)
	route.POST("/", customerController.AddCustomer)
	route.GET("/:id", customerController.GetCustomer)
	route.DELETE("/:id", customerController.DeleteCustomer)

}
