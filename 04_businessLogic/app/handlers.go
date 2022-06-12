package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krlspj/banking-hex-arch/04_businessLogic/service"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello from handler\n")
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	//	customers := []Customer{
	//		{"Aisha", "Tampaulipas", "11005"},
	//		{"Carles", "Barcelona", "80288"},
	//	}
	//fmt.Fprint(w, customers)

	customers, err := ch.service.GetAllCustomers()
	if err != nil {
		log.Fatal(err) // do a proper error handling. Skipping for now...
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		// XML response
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		// JSON response
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)

	}

}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func CreateCostumer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "request content from post..")
}
