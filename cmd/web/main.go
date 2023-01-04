package main

import (
	"fmt"
	"github.com/eniabiola/awesomeProject/pkg/handler"
	"net/http"
)

const portNumber string = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Printf("Starting application on %s", portNumber)

	_ = http.ListenAndServe(
		portNumber,
		nil,
	)
}
