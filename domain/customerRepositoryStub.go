package domain

import "github.com/ameydev/banking-app/errs"

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, *errs.AppError) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Amey", City: "Pune", Zip: "abc-123", Status: "1"},
		{Id: "1001", Name: "Rob", City: "Pune", Zip: "abc-123", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}
