package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	a_mysql "github.com/krlspj/banking-hex-arch/07_security/internal/account/platform/storage/mysql"
	a_service "github.com/krlspj/banking-hex-arch/07_security/internal/account/service"
	c_inmemory "github.com/krlspj/banking-hex-arch/07_security/internal/customer/platform/storage/inmemory"
	c_mysql "github.com/krlspj/banking-hex-arch/07_security/internal/customer/platform/storage/mysql"
	c_service "github.com/krlspj/banking-hex-arch/07_security/internal/customer/service"
	"github.com/krlspj/banking-hex-arch/07_security/internal/drivers"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not defined...")
	}

}

func Start() {

	sanityCheck()

	mux := mux.NewRouter()

	// wiring customer handlers
	dbClient := drivers.OpenConnection()
	customerRepositoryDB := c_mysql.NewCustomerRepositoryDB(dbClient)
	accountRespositoryDB := a_mysql.NewAccountRepositoryDB(dbClient)

	ch := CustomerHandlers{
		service:  c_service.NewCustomerService(customerRepositoryDB),
		servStub: c_service.NewCustomerService(c_inmemory.NewCustomerRepositoryStub()),
	}

	ah := AccountHandlers{
		service: a_service.NewAccountService(accountRespositoryDB),
	}

	// routes
	mux.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	mux.HandleFunc("/customers-mem", ch.getAllCustomersMem).Methods(http.MethodGet)

	mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers", CreateCostumer).Methods(http.MethodPost)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)                                  // make request match -> positives numbers from 0 to 9
	mux.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.CreateAccount).Methods(http.MethodPost)                       // make request match -> positives numbers from 0 to 9
	mux.HandleFunc("/customers-mem/{customer_id:[0-9]+}", ch.getCustomerMem).Methods(http.MethodGet)                           // make request match -> positives numbers from 0 to 9
	mux.HandleFunc("/customers1/{customer_id:[0-9]+}", GetCustomer).Methods(http.MethodGet)                                    // make request match -> positives numbers from 0 to 9
	mux.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost) // make request match -> positives numbers from 0 to 9

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), mux))
}
