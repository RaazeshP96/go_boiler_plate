package repository

import (
	"fmt"
	"log"
	"prototype2/domain"
	"prototype2/errors"

	"github.com/jinzhu/gorm"
)

type productRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{
		DB: db,
	}
}
func (p *productRepository) Save(product *domain.Product) (*domain.Product, error) {
	log.Print("[ProductRepositiory]...Save")
	fmt.Print("[ProductRepositiory]...Save")
	result := p.DB.Create(&product)
	if result.Error != nil {
		err := result.Error
		msg := "error saving the customer"
		err = errors.InternalError.Wrapf(err, msg)
		return nil, err
	}
	return product, nil
}
func (p *productRepository) FindAll() ([]domain.Product, error) {
	log.Print("[ProductRepository]...FindAll")
	fmt.Print("[ProductRepository]...FindAll")
	var product []domain.Product
	result := p.DB.Find(&product)
	if result.Error != nil {
		err := result.Error
		msg := "error getting products"
		err = errors.InternalError.Wrapf(err, msg)
		return nil, err
	}
	return product, nil
}

func (p *productRepository) FindById(id int64) (*domain.Product, error) {
	log.Print("[ProductRepository]...FindById")
	fmt.Print("[ProductRepository]...FindById")
	var product domain.Product
	result := p.DB.Where("id=?", id).First(&product)
	if result.Error != nil {
		err := result.Error
		msg := "error getting the customer with id %d ,id"
		switch err {
		case gorm.ErrRecordNotFound:
			err = errors.NotFound.Wrapf(err, msg)
		default:
			err = errors.InternalError.Wrapf(err, msg)
		}
		return nil, err
	}
	return &product, nil
}

func (p *productRepository) Delete(product *domain.Product) error {
	log.Print("[ProductRepository]...DeleteProduct")
	fmt.Print("[ProductRepository]...DeleteProduct")
	result := p.DB.Delete(product)
	if result.Error != nil {
		err := result.Error
		msg := "error in deleting product"
		err = errors.InternalError.Wrapf(err, msg)
		return err
	}
	return nil
}

func (p *productRepository) Migrate() error {
	log.Print("[ProductRepository]...Migrating")
	fmt.Print("[ProductRepository]...Migrating")
	result := p.DB.AutoMigrate(&domain.Product{})
	switch result.Error {
	case nil:
		return nil
	default:
		return result.Error
	}
}
