package service

import (
	"errors"

	"github.com/hnifmaghfur/bookingtogo/internal/domain"
	"github.com/hnifmaghfur/bookingtogo/internal/interfaces"
)

type customerService struct {
	customerRepo interfaces.CustomerRepository // Dependensi pada antarmuka CustomerRepository
}

func NewCustomerService(customerRepo interfaces.CustomerRepository) interfaces.CustomerService {
	return &customerService{customerRepo: customerRepo}
}

func (s *customerService) CreateCustomer(customer *domain.Customer) error {
	// Contoh logika bisnis: validasi data dasar
	if customer.CstName == "" || customer.CstEmail == "" {
		return errors.New("nama dan email tidak boleh kosong")
	}
	return s.customerRepo.CreateCustomer(customer)
}

func (s *customerService) GetCustomerByID(id uint) (*domain.Customer, error) {
	return s.customerRepo.GetCustomerByID(id)
}

func (s *customerService) GetAllCustomers() ([]domain.Customer, error) {
	return s.customerRepo.GetAllCustomers()
}

func (s *customerService) UpdateCustomer(customer *domain.Customer) error {
	// Contoh logika bisnis: validasi data dasar
	if customer.CstName == "" || customer.CstEmail == "" {
		return errors.New("nama dan email tidak boleh kosong")
	}
	return s.customerRepo.UpdateCustomer(customer)
}

func (s *customerService) DeleteCustomer(id uint) error {
	return s.customerRepo.DeleteCustomer(id)
}
