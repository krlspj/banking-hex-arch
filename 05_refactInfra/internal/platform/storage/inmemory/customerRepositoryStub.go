package inmemory

import "github.com/krlspj/banking-hex-arch/05_refactInfra/internal/domain"

// CustomerRepostoryStub would be the addapter, and should implement
// the Port -> interface from CustomerRespository
type CustomerRepositoryStub struct {
	customers []domain.Customer
}

func NewCustomerRepositoryStub() *CustomerRepositoryStub {
	customers := []domain.Customer{
		{ID: "1000", Name: "Jennipher", City: "Stockolm", Zipcode: "123450", DateOfBirth: "2001-04-26", Status: "1"},
		{ID: "1000", Name: "Natthaya", City: "Stockolm", Zipcode: "123450", DateOfBirth: "1997-10-08", Status: "1"},
	}
	return &CustomerRepositoryStub{
		customers: customers,
	}

}

func (s *CustomerRepositoryStub) FindAll() ([]domain.Customer, error) {
	return s.customers, nil
}
