package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloPricing)
	http.ListenAndServe(":8080", nil)
}

type Response struct {
	Message         string      `json:"message"`
	UpStreamMessage interface{} `json:"up_stream_message"`
}

func HelloPricing(w http.ResponseWriter, r *http.Request) {
	resp := callJsonTest()
	// json.Unmarshal(,)

	var mp map[string]interface{}
	err := json.Unmarshal(resp, &mp)
	if err != nil {
		log.Fatalln(err)
	}

	jsonOut, _ := json.Marshal(
		Response{
			Message:         "Pricing service",
			UpStreamMessage: mp,
		},
	)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", jsonOut)
}

func callJsonTest() []byte {
	resp, err := http.Get("http://time.jsontest.com")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}
