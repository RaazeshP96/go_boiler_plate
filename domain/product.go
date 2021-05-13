package domain

import (
	"time"
)

type Product struct {
	ID        int64     `gorm:"primaryKey autoIncrement"`
	Name      string    `json:"product"`
	Price     float64   `json:"price"`
	Country   string    `json:"country"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type ProductService interface {
	Validate(product *Product) error
	Create(product *Product) (*Product, error)
	FindAll() ([]Product, error)
	GetByID(id string) (*Product, error)
	Delete(id string) error
}

type ProductRepository interface {
	Save(product *Product) (*Product, error)
	FindAll() ([]Product, error)
	FindById(id int64) (*Product, error)
	Delete(product *Product) error
	Migrate() error
}
