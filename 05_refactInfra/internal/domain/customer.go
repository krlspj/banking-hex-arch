package domain

import (
	"github.com/krlspj/banking-hex-arch/05_refactInfra/internal/dto"
	"github.com/krlspj/banking-hex-arch/05_refactInfra/internal/errs"
)

type Customer struct {
	ID          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		ID:          c.ID,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}

// CustomerRepository corresponds to the port of Curstmer
// contains the contracts for Costumer
type CustomerRepository interface {
	// status == 1 status == 0 status == ""
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
