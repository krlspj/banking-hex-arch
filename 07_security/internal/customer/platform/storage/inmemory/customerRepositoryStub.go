package inmemory

import (
	"github.com/krlspj/banking-hex-arch/07_security/internal/customer/domain"
	"github.com/krlspj/banking-hex-arch/07_security/internal/errs"
)

// CustomerRepostoryStub would be the addapter, and should implement
// the Port -> interface from CustomerRespository
type CustomerRepositoryStub struct {
	customers []domain.Customer
}

func NewCustomerRepositoryStub() *CustomerRepositoryStub {
	customers := []domain.Customer{
		{ID: "1000", Name: "Jennipher", City: "Stockolm", Zipcode: "123450", DateOfBirth: "2001-04-26", Status: "1"},
		{ID: "1001", Name: "Natthaya", City: "Stockolm", Zipcode: "123450", DateOfBirth: "1997-10-08", Status: "0"},
	}
	return &CustomerRepositoryStub{
		customers: customers,
	}

}

func (s *CustomerRepositoryStub) FindAll(status string) ([]domain.Customer, *errs.AppError) {
	if status == "" {
		return s.customers, nil

	} else if status == "1" || status == "0" {
		filtredCustomers := []domain.Customer{}
		for _, v := range s.customers {
			if v.Status == status {
				filtredCustomers = append(filtredCustomers, v)
			}
		}
		return filtredCustomers, nil
	} else {
		return nil, nil
	}
}

func (s *CustomerRepositoryStub) ById(id string) (*domain.Customer, *errs.AppError) {
	for _, v := range s.customers {
		if v.ID == id {
			return &v, nil
		}
	}
	//return nil, errors.New("User Not found")
	return nil, errs.NewNotFoundError("User Not found")
}
