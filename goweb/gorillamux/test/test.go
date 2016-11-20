package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type fruit struct {
	ID       int    `json:"id,omitempty"`
	Kind     string `json:"kind,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}

var basket []fruit

func main() {

	f1 := fruit{1, "banana", 10}
	f2 := fruit{2, "apple", 5}
	basket = append(basket, f1)
	basket = append(basket, f2)
	r := mux.NewRouter()
	r.HandleFunc("/fruits", fruits).Methods("GET")
	log.Fatal(http.ListenAndServe(":9000", r))
}

func fruits(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(basket)
}
