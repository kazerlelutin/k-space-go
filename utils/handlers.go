package utils

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hello from Koyeb")
}

func MyFirstEndpoint(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "This is my first endpoint -->")
}
