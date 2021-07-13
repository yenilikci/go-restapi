package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	. "../helpers"
	. "../models"
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

}

//HTTP Put - /api/products/{id}
func PutProductHandler(w http.ResponseWriter, r *http.Request) {

}

//HTTP Delete - /api/products/{id}
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

}
