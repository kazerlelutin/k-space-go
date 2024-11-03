package main

import (
	"fmt"
	"k-space-go/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", utils.HelloHandler)
	http.HandleFunc("/my-first-endpoint", utils.MyFirstEndpoint)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
