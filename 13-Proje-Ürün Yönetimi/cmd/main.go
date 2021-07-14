package main

import (
	"log"
	"net/http"

	. "../handlers"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("server starting...")
	r := mux.NewRouter().StrictSlash(true) // end /
	r.HandleFunc("/api/products", GetProductsHandler).Methods("GET")
	r.HandleFunc("/api/products/{id}", GetProductHandler).Methods("GET")
	r.HandleFunc("/api/products", PostProductHandler).Methods("POST")
	r.HandleFunc("/api/products/{id}", PutProductHandler).Methods("PUT")
	r.HandleFunc("/api/products/{id}", DeleteProductHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	server.ListenAndServe()
	log.Println("server ending...")
}
