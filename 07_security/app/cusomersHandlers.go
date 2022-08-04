package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krlspj/banking-hex-arch/07_security/internal/customer/service"
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

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllCustomers")
	//	customers := []Customer{
	//		{"Aisha", "Tampaulipas", "11005"},
	//		{"Carles", "Barcelona", "80288"},
	//	}
	//fmt.Fprint(w, customers)
	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomers(status)

	if err != nil {
		log.Println(err) // do a proper error handling. Skipping for now...
		writeResponse(w, err.Code, err)
		return

	}
	//remap all customers responses
	//respCustomers := make([]Customer, 0, len(customers))
	//for _, v := range customers {
	//	customer := Customer{
	//		Name:    v.Name,
	//		City:    v.City,
	//		Zipcode: v.Zipcode,
	//	}
	//	respCustomers = append(respCustomers, customer)
	//}

	fmt.Println("respCustomers =>", customers)
	if r.Header.Get("Content-Type") == "application/xml" {
		// XML response
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		// JSON response
		//writeResponse(w, http.StatusOK, customers)
		writeResponse(w, http.StatusOK, customers)

	}

}
func (ch *CustomerHandlers) getAllCustomersMem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllCustomers")
	//	customers := []Customer{
	//		{"Aisha", "Tampaulipas", "11005"},
	//		{"Carles", "Barcelona", "80288"},
	//	}
	//fmt.Fprint(w, customers)
	status := r.URL.Query().Get("status")
	customers, err := ch.servStub.GetAllCustomers(status)

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
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(err.Code)
		//fmt.Fprintf(w, err.Message)

		json.NewEncoder(w).Encode(err)
		//json.NewEncoder(w).Encode(err.AsMessage())
	} else {
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customer)
	}
}
func (ch *CustomerHandlers) getCustomerMem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.servStub.GetCustomer(id)
	if err != nil {
		//fmt.Fprintf(w, err.Message)
		writeResponse(w, err.Code, err)
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}
func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func CreateCostumer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "request content from post..")
}

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
