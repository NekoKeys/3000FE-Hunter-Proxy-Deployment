package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ProxyRequest struct {
	URL string
}

func proxyRequest(w http.ResponseWriter, r *http.Request) {
	var p ProxyRequest

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := &http.Client{}

	req, err := http.NewRequest("GET", p.URL, nil)
	req.Header = r.Header

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Fprintf(w, "%s", string(body))
}

func handleRequests() {
	http.HandleFunc("/proxy-request", proxyRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
