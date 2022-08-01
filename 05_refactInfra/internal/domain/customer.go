package domain

import "github.com/krlspj/banking-hex-arch/05_refactInfra/internal/errs"

type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

// CustomerRepository corresponds to the port of Curstmer
// contains the contracts for Costumer
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
