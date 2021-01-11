package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Customer struct {
	Name    string
	Age     int
	Address string
}

// Test directly calls handler func to test for outputs.
func handler(w http.ResponseWriter, r *http.Request) {
	// reads bytes from input output.
	body, _ := ioutil.ReadAll(r.Body)
	// read the bytes, after converting them into type interface/ struct here, using unmarshal.
	// Create an output var customer
	var c Customer
	err := json.Unmarshal(body, &c)
	if err != nil {
		log.Fatalf("%s", err)
	}

	//fmt.Println(c)
	if c.Age == 0 {
		io.WriteString(w, "not eligible")
	} else if c.Age < 18 {
		io.WriteString(w, "not eligible")
	} else {
		// Can use \ (backslash) to ignore " in fmt.sprintf for formatting.
		io.WriteString(w, fmt.Sprintf("{\"Name\" : \"%s\", \"Age\" : %v, \"Address\" : \"%s\"}", c.Name, c.Age, c.Address))
	}
}

// Main spawns the server.
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// For code coverage:
// go test ./... -v -coverprofile coverage.txt -coverpkg=./...
// go tool cover -html coverage.txt
