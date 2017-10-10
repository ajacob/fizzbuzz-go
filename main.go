package main

import (
	"log"
	"net/http"
	"strconv"
)

const port int = 1337

// Entry point of the application.
// This main function will simply register a fizzBuzzHandler and start
// serving requests.
func main() {
	http.HandleFunc("/", fizzBuzzHandler)

	log.Println("Starting server on port", port)

	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}
