package interfaces

import (
	"github.com/hnifmaghfur/bookingtogo/internal/domain"
)

type CustomerRepository interface {
	CreateCustomer(customer *domain.Customer) error
	GetCustomerByID(cstId uint) (*domain.Customer, error)
	GetAllCustomers() ([]domain.Customer, error)
	UpdateCustomer(customer *domain.Customer) error
	DeleteCustomer(cstId uint) error
}

type CustomerService interface {
	CreateCustomer(customer *domain.Customer) error
	GetCustomerByID(cstId uint) (*domain.Customer, error)
	GetAllCustomers() ([]domain.Customer, error)
	UpdateCustomer(customer *domain.Customer) error
	DeleteCustomer(cstId uint) error
}
