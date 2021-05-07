package domain

type Product struct {
	ID    int64   `gorm:"primaryKey autoIncrement"`
	Price float64 `json:"price"`
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
