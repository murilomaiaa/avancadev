package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Result struct {
	Status string
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9093", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello microservices")

	result := Result{Status: "ok"}

	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Fatal("Error processing JSON")
	}

	fmt.Fprintf(w, string(jsonData))
}
