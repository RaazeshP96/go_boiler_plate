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

type customerController struct {
	customerService domain.CustomerService
}
type CustomerController interface {
	GetCustomers(c *gin.Context)
	AddCustomer(c *gin.Context)
	GetCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
}

func NewCustomerController(cus domain.CustomerService) CustomerController {
	return &customerController{
		customerService: cus,
	}
}

func (cus *customerController) GetCustomers(c *gin.Context) {
	log.Printf("[Customer]...GetCustomers")
	fmt.Printf("[Customer]...GetCustomers")
	customers, err := cus.customerService.FindAll()
	if err != nil {
		responses.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (cus *customerController) AddCustomer(c *gin.Context) {
	log.Print("[CustomerControllere]...AddCustomer")
	fmt.Print("[CustomerControllere]...AddCustomer")
	var customer domain.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		err = errors.BadRequest.New("error parsing the input information")
		responses.HandleError(c, err)
	}

	// customer.ID = rand.Int63()
	if err := cus.customerService.Validate(&customer); err != nil {
		responses.HandleError(c, err)
		return
	}
	customerCreated, err := cus.customerService.Create(&customer)
	if err != nil {
		responses.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, customerCreated)
}

func (cus *customerController) GetCustomer(c *gin.Context) {
	log.Print("[CustomerController]...GetCustomer")
	customer, err := cus.customerService.GetByID(c.Param("id"))
	if err != nil {
		responses.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": customer})
}

func (cus *customerController) DeleteCustomer(c *gin.Context) {
	log.Print("[CustomerController]...DeleteCustomer")
	err := cus.customerService.Delete(c.Param("id"))
	if err != nil {
		responses.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": true})
}
