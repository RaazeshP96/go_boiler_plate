package service

import (
	"log"
	"prototype2/domain"
	"prototype2/errors"
	"strconv"
	"sync"
)

var once sync.Once

type customerService struct {
	repo domain.CustomerRepository
}

var instance *customerService

func NewCustomerService(r domain.CustomerRepository) domain.CustomerService {
	once.Do(func() {
		instance = &customerService{
			repo: r,
		}
	})
	return instance
}

func (*customerService) Validate(customer *domain.Customer) error {
	log.Print("[CustomerService]...Validate")
	if customer == nil {
		err := errors.BadRequest.New("The customer is empty")
		return err
	}
	return nil
}

func (c *customerService) Create(customer *domain.Customer) (*domain.Customer, error) {
	log.Print("[CustomerService]...Create")
	// customer.ID = rand.Int63()
	return c.repo.Save(customer)
}

func (c *customerService) FindAll() ([]domain.Customer, error) {
	log.Print("[CustomerService]...FindAll")
	return c.repo.FindAll()
}

func (c *customerService) GetByID(idString string) (*domain.Customer, error) {
	log.Print("[CustomerService]...GetById")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		err = errors.BadRequest.Wrapf(err, "interactor converting id to int")
		err = errors.AddErrorContext(err, "id", "wrong id format")
		return nil, err
	}
	return c.repo.FindByID(id)
}

func (c *customerService) Delete(idString string) error {
	log.Print("[CustomerService]...Delete")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		err = errors.BadRequest.Wrapf(err, "interactor converting id to int")
		err = errors.AddErrorContext(err, "id", "wrong id format")
		return err
	}
	customer, err := c.repo.FindByID(id)
	if err != nil {
		return err
	}
	return c.repo.Delete(customer)
}
