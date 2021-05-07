package domain

type Product struct {
	ID    int64 `gorm:"primaryKey autoIncrement"`
	Price int64 `json:"price"`
}

type ProductRepository interface {
	Save(product *Product) (*Product, error)
	FindAll() ([]Product, error)
	FindById(id int64) (*Product, error)
	Delete(product *Product) error
	Migrate() error
}

type ProductService interface {
	validate(product *Product) error
	Create(product *Product) (*Product, error)
	FindAll() ([]Product, error)
	GetByID(id int64) (*Product, error)
	Delete(id int64) error
}
