package service

import (
	"fmt"
	"log"
	"prototype2/domain"
	"prototype2/errors"
	"strconv"
	"sync"
)

var once sync.Once

type productService struct {
	repo domain.ProductRepository
}

var instance *productService

func NewProductService(r domain.ProductRepository) domain.ProductService {
	once.Do(func() {
		instance = &productService{
			repo: r,
		}
	})
	return instance
}

func (p *productService) Validate(product *domain.Product) error {
	log.Print("[ProductService]...Validate")
	fmt.Print("[ProductService]...Validate")
	if product == nil {
		err := errors.BadRequest.New("The product is empty")
		return err
	}
	return nil
}

func (p *productService) Create(product *domain.Product) (*domain.Product, error) {
	log.Print("[ProductService]...ProductCreate")
	fmt.Print("[ProductService]...ProductCreate")
	return p.repo.Save(product)
}

func (p *productService) FindAll() ([]domain.Product, error) {
	log.Print("[ProductService]...FindAll")
	fmt.Print("[ProductService]...FindAll")
	return p.repo.FindAll()
}

func (p *productService) GetByID(idString string) (*domain.Product, error) {
	log.Print("[ProductService]...FindById")
	fmt.Print("[ProductService]...FindById")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		err = errors.BadRequest.Wrapf(err, "interactor converting id to int")
		err = errors.AddErrorContext(err, "id", "wrong id format")
		return nil, err
	}
	return p.repo.FindById(id)
}

func (p *productService) Delete(idString string) error {
	log.Print("[ProductService]...Delete")
	fmt.Print("[ProductService]...Delete")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		err = errors.BadRequest.Wrapf(err, "interactor converting id to int")
		err = errors.AddErrorContext(err, "id", "wrong id format")
		return err
	}
	product, err := p.repo.FindById(id)
	if err != nil {
		return err
	}
	return p.repo.Delete(product)
}
