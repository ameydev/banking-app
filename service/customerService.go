package service

import (
	"github.com/ameydev/banking-app/domain"
	"github.com/ameydev/banking-app/errs"
)

// it is the primary port
type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, *errs.AppError)
	GetCustomersByStatus(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

// Adaptor
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomersByStatus(status string) ([]domain.Customer, *errs.AppError) {
	return s.repo.ByStatus(status)
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	// this is the implementation of the service where dependency of the secondary port is injected
	return DefaultCustomerService{repo: repository}
}
