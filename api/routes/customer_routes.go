package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func PostRoutes(route *gin.RouterGroup, db *gorm.DB) {
	CustomerRepository := customer_repository.NewCustomerRepository(db)
	if err := postRepository.Migrate(); err != nil {
		log.Fatal("Customer migrate err", err)
	}
	customerService := customer_service.NewCustomerService(customerRepository)
	customerController := customer_controller.NewPostController(customerService)

	route.GET("/", customerControllerr.GetCustomers)
	route.POST("/", customerControllerr.AddCustomer)
	route.GET("/:id", customerControllerr.GetCustomer)
	route.DELETE("/:id", customerControllerr.DeleteCustomers)

}
