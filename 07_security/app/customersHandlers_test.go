package app

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/krlspj/banking-hex-arch/07_security/internal/customer/dto"
	"github.com/krlspj/banking-hex-arch/07_security/internal/customer/service/servicemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetCustomersStatusOK(t *testing.T) {
	//ctrl := mock.NewCustomerService(t)

	mockService := new(servicemocks.CustomerService)
	//mockService := servicemocks.NewCustomerService(t)

	ch := CustomerHandlers{service: mockService}
	//ch.service.GetAllCustomers("active")

	dummyCustomers := []dto.CustomerResponse{
		{Name: "Aisha", City: "Tampaulipas", Zipcode: "11005", Status: "active"},
		{Name: "Carles", City: "Barcelona", Zipcode: "80288", Status: "inactive"},
	}

	mockService.On("GetAllCustomers", mock.AnythingOfType("string")).Return(dummyCustomers, nil)

	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)
	request, err := http.NewRequest(http.MethodGet, "/customers", nil)
	fmt.Println("----- request", request)
	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	//
	if recorder.Code != http.StatusOK {
		t.Error("Failde while testing status code")
	}

}
