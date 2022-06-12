package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krlspj/banking-hex-arch/04_businessLogic/domain"
	"github.com/krlspj/banking-hex-arch/04_businessLogic/service"
)

func Start() {

	mux := mux.NewRouter()

	// wiring customer handlers
	ch := CustomerHandlers{
		//		service.NewCustomerService(domain.NewCustomerRepositoryStub()),
		service: service.NewCustomerService(domain.NewCustomerRepositoryDB()),
	}

	// routes
	mux.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)

	mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers", CreateCostumer).Methods(http.MethodPost)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", GetCustomer).Methods(http.MethodGet) // make request match -> positives numbers from 0 to 9

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
