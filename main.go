package main

import (
	"log"
	"net/http"
)

// Entry point of the application.
// This main function will simply register a fizzBuzzHandler and start
// serving requests on port 1337.
func main() {
	http.HandleFunc("/", fizzBuzzHandler)

	log.Fatal(http.ListenAndServe(":1337", nil))
}
