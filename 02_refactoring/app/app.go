package app

import (
	"log"
	"net/http"
)

func Start() {

	mux := http.NewServeMux()
	//	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
	//		fmt.Fprint(w, "hello from handler\n")
	//	})
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", GetAllCostumers)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
