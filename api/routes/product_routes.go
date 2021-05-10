package routes

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	product_controller "prototype2/api/controller/product"
	product_repository "prototype2/api/repository/product"
	product_service "prototype2/api/service/product"
)

func ProductRoutes(route *gin.RouterGroup, db *gorm.DB) {
	productRepository := product_repository.NewProductRepository(db)
	if err := productRepository.Migrate(); err != nil {
		log.Fatal("Product migrate err", err)
		fmt.Printf("Product migrate err")
	}
	productService := product_service.NewProductService(productRepository)
	productController := product_controller.NewProductController(productService)

	route.GET("/", productController.GetProducts)
	route.POST("/", productController.AddProduct)
	route.GET("/:id", productController.GetProduct)
	route.DELETE("/:id", productController.DeleteProduct)

}
