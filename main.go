package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/my-first-endpoint", myFirstEndpoint)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello from Koyeb")
}

func myFirstEndpoint(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "This is my first endpoint -->")
}
