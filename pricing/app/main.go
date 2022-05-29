package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloProduct)
	http.HandleFunc("/pricing", HelloPricing)
	http.ListenAndServe(":8080", nil)
}

type Response struct {
	Message string `json:"message"`
}

func HelloProduct(w http.ResponseWriter, r *http.Request) {
	jsonOut, _ := json.Marshal(Response{Message: "Hello GUYS?"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jsonOut)
}

func HelloPricing(w http.ResponseWriter, r *http.Request) {
	jsonOut, _ := json.Marshal(Response{Message: "Hello Pricing?"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jsonOut)
}
