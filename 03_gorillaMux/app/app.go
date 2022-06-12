package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	//mux := http.NewServeMux()
	mux := mux.NewRouter()
	//	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
	//		fmt.Fprint(w, "hello from handler\n")
	//	})
	mux.HandleFunc("/greet", greet).Methods(http.MethodGet)
	mux.HandleFunc("/customers", GetAllCostumers).Methods(http.MethodGet)
	mux.HandleFunc("/customers", CreateCostumer).Methods(http.MethodPost)
	mux.HandleFunc("/customers/{customer_id:[0-9]+}", GetCustomer).Methods(http.MethodGet) // make request match -> positives numbers from 0 to 9

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
