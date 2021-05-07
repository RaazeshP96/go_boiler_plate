package repository

import (
	"log"
	"prototype2/domain"

	"prototype2/errors"

	"github.com/jinzhu/gorm"
)

type customerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) domain.CustomerRepository {
	return &customerRepository{
		DB: db,
	}
}

func (c *customerRepository) Save(customer *domain.Customer) (*domain.Customer, error) {
	log.Print("[CustomerRepository]...Save")
	result := c.DB.Create(&customer)
	if result.Error != nil {
		err := result.Error
		msg := "error saving the customer"
		err = errors.InternalError.Wrapf(err, msg)
		return nil, err
	}
	return customer, nil

}

func (c *customerRepository) FindAll() ([]domain.Customer, error) {
	log.Print("[CustomerRepository]...FindAll")
	var customers []domain.Customer
	result := c.DB.Find(&customers)
	if result.Error != nil {
		err := result.Error
		msg := "error getting the customers"
		switch err {
		case gorm.ErrRecordNotFound:
			err = errors.NotFound.Wrapf(err, msg)
		default:
			err = errors.InternalError.Wrapf(err, msg)
		}
		return nil, err
	}
	return customers, nil
}

func (c *customerRepository) FindByID(id int64) (*domain.Customer, error) {
	log.Print("[CustomerRepository]...FindById")
	var customer domain.Customer
	result := c.DB.Where("id=?", id).First(&customer)
	if result.Error != nil {
		err := result.Error
		msg := "error getting the customer  with id %d,id"
		switch err {
		case gorm.ErrRecordNotFound:
			err = errors.NotFound.Wrapf(err, msg)
		default:
			err = errors.InternalError.Wrapf(err, msg)
		}
		return nil, err
	}
	return &customer, nil
}

func (c *customerRepository) Delete(customer *domain.Customer) error {
	log.Print("[CustomerRepository]...Delete")
	result := c.DB.Delete(&customer)
	if result.Error != nil {
		err := result.Error
		msg := "error in deleting post"
		err = errors.InternalError.Wrapf(err, msg)
		return err
	}
	return nil
}

func (c *customerRepository) Migrate() error {
	log.Print("[CustomerRepository]...Migrate")
	result := c.DB.AutoMigrate(&domain.Customer{})

	switch result.Error {
	case nil:
		return nil
	default:
		return result.Error
	}
}
