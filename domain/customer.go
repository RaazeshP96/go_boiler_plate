package domain

type Customer struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CustomerService interface {
	Validate(customer *Customer) error
	Create(customer *Customer) (*Customer, error)
	FindAll() ([]Customer, error)
	GetByID(id string) (*Customer, error)
	Delete(id string) error
}

type CustomerRepository interface {
	Save(customer *Customer) (*Customer, error)
	FindAll() ([]Customer, error)
	FindByID(id int64) (*Customer, error)
	Delete(customer *Customer) error
	Migrate() error
}
