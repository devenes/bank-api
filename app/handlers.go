package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/devenes/bank-api/service"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode int    `json:"zipcode" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService // dependency injection

}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	/*	customers := []Customer{
		Customer{
			Name:    "John",
			City:    "New York",
			Zipcode: 10001,
		},
		Customer{
			Name:    "Jane",
			City:    "New York",
			Zipcode: 10001,
		},
	}*/
	customers, err := ch.service.GetAllCustomer()
	if err != nil {
		// handle the error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
