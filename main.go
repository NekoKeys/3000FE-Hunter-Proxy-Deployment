package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ProxyRequest struct {
	URL string
}

func proxyRequest(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	var p ProxyRequest

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	duration := time.Since(start)
	fmt.Println(duration)

	client := &http.Client{}

	req, err := http.NewRequest("GET", p.URL, nil)
	req.Header = r.Header

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	duration = time.Since(start)
	fmt.Println(duration)

	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)

	duration = time.Since(start)
	fmt.Println(duration)

	fmt.Fprintf(w, "%s", bodyString)

	fmt.Println("Endpoint Hit: homePage")
	duration = time.Since(start)
	fmt.Println(duration)

}

func handleRequests() {
	http.HandleFunc("/proxy-request", proxyRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
