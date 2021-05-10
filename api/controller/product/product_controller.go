package controller

import (
	"fmt"
	"log"
	"net/http"
	"prototype2/api/responses"
	"prototype2/domain"
	"prototype2/errors"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productService domain.ProductService
}

type ProductController interface {
	GetProducts(c *gin.Context)
	AddProduct(c *gin.Context)
	GetProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

func NewProductController(p domain.ProductService) ProductController {
	return &productController{
		productService: p,
	}
}

func (p *productController) GetProducts(c *gin.Context) {
	log.Printf("[Product]...GetCustomers")
	fmt.Printf("[Product]...GetCustomers")
	products, err := p.productService.FindAll()
	if err != nil {
		responses.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, products)
}

func (p *productController) AddProduct(c *gin.Context) {
	log.Printf("[ProductController]...AddProduct")
	fmt.Printf("[ProductController]...AddProduct")
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		err = errors.BadRequest.New("error parsing the input information")
		responses.HandleError(c, err)
	}
	productCreated, err := p.productService.Create(&product)
	if err != nil {
		responses.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, productCreated)
}

func (p *productController) GetProduct(c *gin.Context) {
	log.Printf("[ProductController]...GetProduct")
	fmt.Printf("[ProductController]...GetProduct")
	customer, err := p.productService.GetByID(c.Param("id"))
	if err != nil {
		responses.HandleError(c, err)
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}
func (p *productController) DeleteProduct(c *gin.Context) {
	log.Printf("[ProductController]...DeleteProduct")
	fmt.Printf("[ProductController]...DeleteProduct")
	err := p.productService.Delete(c.Param("id"))
	if err != nil {
		responses.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
