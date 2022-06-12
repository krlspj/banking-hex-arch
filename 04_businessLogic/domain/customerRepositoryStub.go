package domain

// CustomerRepostoryStub would be the addapter, and should implement
// the Port -> interface from CustomerRespository
type CustomerRepositoryStub struct {
	customers []Customer
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1000", "Jennipher", "Stockolm", "123450", "2001-04-26", "1"},
		{"1000", "Natthaya", "Stockolm", "123450", "1997-10-08", "1"},
	}
	return CustomerRepositoryStub{
		customers: customers,
	}

}

func (s *CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}
