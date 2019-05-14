package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// get configuration
	address := flag.String("server", "http://localhost:8080", "HTTP gateway url, e.g. http://localhost:8080")
	flag.Parse()

	//t := time.Now().In(time.UTC)
	//pfx := t.Format(time.RFC3339Nano)

	var body string

	type reg struct {
		name string `json:"name"`
	}

	// Call Read
	resp, err := http.Get(fmt.Sprintf("%s%s%s", *address, "/v1/greet", ""))
	if err != nil {
		log.Fatalf("failed to call Read method: %v", err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		body = fmt.Sprintf("failed read Read response body: %v", err)
	} else {
		body = string(bodyBytes)
	}
	log.Printf("Read response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

}
