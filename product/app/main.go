package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HelloProduct)
	http.ListenAndServe(":8080", nil)
}

type Response struct {
	Message         string      `json:"message"`
	UpStreamMessage interface{} `json:"up_stream_message"`
}

func HelloProduct(w http.ResponseWriter, r *http.Request) {
	resp := callPricingService()
	var mp map[string]interface{}
	err := json.Unmarshal(resp, &mp)
	if err != nil {
		log.Fatalln(err)
	}
	jsonOut, _ := json.Marshal(
		Response{
			Message:         "Product service",
			UpStreamMessage: mp,
		},
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jsonOut)
}

func callPricingService() []byte {
	resp, err := http.Get(os.Getenv("PRICING_SVC_URI"))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}
