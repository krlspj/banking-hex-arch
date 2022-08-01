package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krlspj/banking-hex-arch/05_refactInfra/internal/service"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

type CustomerHandlers struct {
	service  service.CustomerService
	servStub service.CustomerService
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello from handler\n")
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllCustomers")
	//	customers := []Customer{
	//		{"Aisha", "Tampaulipas", "11005"},
	//		{"Carles", "Barcelona", "80288"},
	//	}
	//fmt.Fprint(w, customers)

	customers, err := ch.service.GetAllCustomers()

	if err != nil {
		log.Fatal(err) // do a proper error handling. Skipping for now...
	}
	respCustomers := make([]Customer, 0, len(customers))
	for _, v := range customers {
		customer := Customer{
			Name:    v.Name,
			City:    v.City,
			Zipcode: v.Zipcode,
		}
		respCustomers = append(respCustomers, customer)
	}
	fmt.Println("respCustomers =>", respCustomers)
	if r.Header.Get("Content-Type") == "application/xml" {
		// XML response
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(respCustomers)
	} else {
		// JSON response
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(respCustomers)

	}

}
func (ch *CustomerHandlers) GetAllCustomersMem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllCustomers")
	//	customers := []Customer{
	//		{"Aisha", "Tampaulipas", "11005"},
	//		{"Carles", "Barcelona", "80288"},
	//	}
	//fmt.Fprint(w, customers)

	customers, err := ch.servStub.GetAllCustomers()

	if err != nil {
		log.Fatal(err) // do a proper error handling. Skipping for now...
	}
	respCustomers := make([]Customer, 0, len(customers))
	for _, v := range customers {
		customer := Customer{
			Name:    v.Name,
			City:    v.City,
			Zipcode: v.Zipcode,
		}
		respCustomers = append(respCustomers, customer)
	}
	fmt.Println("respCustomers =>", respCustomers)
	if r.Header.Get("Content-Type") == "application/xml" {
		// XML response
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(respCustomers)
	} else {
		// JSON response
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(respCustomers)

	}

}
func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
func (ch *CustomerHandlers) getCustomerMem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.servStub.GetCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func CreateCostumer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "request content from post..")
}
