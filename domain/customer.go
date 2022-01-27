package domain

import "github.com/ameydev/banking-app/errs"

type Customer struct {
	Id          string
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zip         string `json:"zip"`
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
	ByStatus(string) ([]Customer, *errs.AppError)
}
