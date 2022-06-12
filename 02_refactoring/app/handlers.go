package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello from handler\n")
}

func GetAllCostumers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Aisha", "Tampaulipas", "11005"},
		{"Carles", "Barcelona", "80288"},
	}
	//fmt.Fprint(w, customers)

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
