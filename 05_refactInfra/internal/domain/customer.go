package domain

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
	ById(string) (*Customer, error)
}
