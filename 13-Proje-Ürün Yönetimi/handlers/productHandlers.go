package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	. "../helpers"
	. "../models"
	"github.com/gorilla/mux"
)

var productStore = make(map[string]Product)
var id int = 0

//HTTP Post - /api/products
func PostProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	CheckError(err)

	product.CreatedOn = time.Now()
	id++
	product.ID = id
	key := strconv.Itoa(id)
	productStore[key] = product

	//go obj -> json
	data, err := json.Marshal(product)
	CheckError(err)
	//write
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

//HTTP Get - /api/products
func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	var products []Product
	for _, product := range productStore {
		products = append(products, product)
	}

	//go obj -> json
	data, err := json.Marshal(products)
	CheckError(err)
	//write
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

//HTTP Post - /api/products/{id}
func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	//mux Vars
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	for _, prd := range productStore {
		if prd.ID == key {
			product = prd
		}
	}

	//go obj -> json
	data, err := json.Marshal(product)
	CheckError(err)
	//write
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

//HTTP Put - /api/products/{id}
func PutProductHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	key := vars["id"]

	var prodUpd Product
	err = json.NewDecoder(r.Body).Decode(&prodUpd)
	CheckError(err)

	if _, ok := productStore[key]; ok {
		prodUpd.ID, _ = strconv.Atoi(key)
		prodUpd.ChangedOn = time.Now()
		delete(productStore, key)
		productStore[key] = prodUpd
	} else {
		log.Printf("Değer bulunamadı: %s", key)
	}
	w.WriteHeader(http.StatusOK)
}

//HTTP Delete - /api/products/{id}
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	if _, ok := productStore[key]; ok {
		delete(productStore, key)
	} else {
		log.Printf("Değer bulunamadı: %s", key)
	}
	w.WriteHeader(http.StatusOK)
}
