package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Starting server on port 8080...")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go! 🎯\n")
}