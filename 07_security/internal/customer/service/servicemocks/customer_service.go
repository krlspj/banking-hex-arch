// Code generated by mockery v2.14.0. DO NOT EDIT.

package servicemocks

import (
	dto "github.com/krlspj/banking-hex-arch/07_security/internal/customer/dto"
	errs "github.com/krlspj/banking-hex-arch/07_security/internal/errs"

	mock "github.com/stretchr/testify/mock"
)

// CustomerService is an autogenerated mock type for the CustomerService type
type CustomerService struct {
	mock.Mock
}

// GetAllCustomers provides a mock function with given fields: _a0
func (_m *CustomerService) GetAllCustomers(_a0 string) ([]dto.CustomerResponse, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 []dto.CustomerResponse
	if rf, ok := ret.Get(0).(func(string) []dto.CustomerResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.CustomerResponse)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(string) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// GetCustomer provides a mock function with given fields: _a0
func (_m *CustomerService) GetCustomer(_a0 string) (*dto.CustomerResponse, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *dto.CustomerResponse
	if rf, ok := ret.Get(0).(func(string) *dto.CustomerResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.CustomerResponse)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(string) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

type mockConstructorTestingTNewCustomerService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCustomerService creates a new instance of CustomerService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCustomerService(t mockConstructorTestingTNewCustomerService) *CustomerService {
	mock := &CustomerService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
