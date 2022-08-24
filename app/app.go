package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// declare a new router
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// Start the server
	log.Fatalf("Server failed to start: %v", http.ListenAndServe("localhost:8080", router))
}

func createCustomer(w http.ResponseWriter, request *http.Request) {
	fmt.Fprint(w, "Post request is received.")
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	_, err := fmt.Fprint(w, vars["customer_id"])
	if err != nil {
		return
	}
}
