package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krlspj/banking-hex-arch/05_refactInfra/internal/platform/storage/inmemory"
	"github.com/krlspj/banking-hex-arch/05_refactInfra/internal/platform/storage/mysql"
	"github.com/krlspj/banking-hex-arch/05_refactInfra/internal/service"
)

func Start() {
	log.Println("Server listening on port 8000")

	mux := mux.NewRouter()

	// wiring customer handlers
	ch := CustomerHandlers{
		service:  service.NewCustomerService(mysql.NewCustomerRepositoryDB()),
		servStub: service.NewCustomerService(inmemory.NewCustomerRepositoryStub()),
	}

	// routes
	mux.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customers-mem", ch.GetAllCustomersMem).Methods(http.MethodGet)

	mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers", CreateCostumer).Methods(http.MethodPost)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)        // make request match -> positives numbers from 0 to 9
	mux.HandleFunc("/customers-mem/{customer_id:[0-9]+}", ch.getCustomerMem).Methods(http.MethodGet) // make request match -> positives numbers from 0 to 9
	mux.HandleFunc("/customers1/{customer_id:[0-9]+}", GetCustomer).Methods(http.MethodGet)          // make request match -> positives numbers from 0 to 9

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
