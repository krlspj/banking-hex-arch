package service

import (
	"github.com/krlspj/banking-hex-arch/05_refactInfra/internal/domain"
	"github.com/krlspj/banking-hex-arch/05_refactInfra/internal/dto"
	"github.com/krlspj/banking-hex-arch/05_refactInfra/internal/errs"
)

type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	cl, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	var resp []dto.CustomerResponse
	for _, v := range cl {
		resp = append(resp, v.ToDto())
	}
	return resp, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {

	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	resp := c.ToDto()

	return &resp, nil
}

func NewCustomerService(repository domain.CustomerRepository) *DefaultCustomerService {
	return &DefaultCustomerService{
		repo: repository,
	}
}
