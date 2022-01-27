package app

import (
	"log"
	"net/http"

	"github.com/ameydev/banking-app/domain"
	"github.com/ameydev/banking-app/service"
	"github.com/gorilla/mux"
)

var dbURL string = "root:codecamp@tcp(localhost:3306)/banking"

func Start() {
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB(dbURL))}
	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", router))
}
