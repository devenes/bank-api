package app

import (
	"fmt"
	"github.com/devenes/bank-api/domain"
	"github.com/devenes/bank-api/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring up the handlers
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// declare a new router
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

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
