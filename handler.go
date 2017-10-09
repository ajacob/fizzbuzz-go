package main

// This class contains the fizzBuzzHandler that serves HTTP requests.
// After having fetched and checked the parameters, it will call the fizz buzz
// algorithm with an implementation that renders the result as JSON.

import (
	"net/http"
	"errors"
	"encoding/json"
	"log"
)

const defaultString1 string = "fizz"
const defaultString2 string = "buzz"

const defaultInt1 int = 3
const defaultInt2 int = 5

const defaultLimit int = 100

// responseError describes an error in a REST context.
// Provides a status code and a message describing the error to the user.
type responseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Returns a parameters struct filled from the request.
// They will be converted to the correct type, checked and defaulted if missing.
func fetchAndCheckParameters(request *http.Request) (*parameters, error) {
	var err error

	params := parameters{}

	params.string1 = defaultString(request.FormValue("string1"), defaultString1)
	params.string2 = defaultString(request.FormValue("string2"), defaultString2)

	if params.int1, err = defaultAtoi(request.FormValue("int1"), defaultInt1); err != nil {
		return nil, errors.New("int1 must be a valid integer")
	}

	if params.int2, err = defaultAtoi(request.FormValue("int2"), defaultInt2); err != nil {
		return nil, errors.New("int2 must be a valid integer")
	}

	// We don't want to divide by zero !
	if params.int1 == 0 || params.int2 == 0 {
		return nil, errors.New("int1 and int2 must not be 0")
	}

	if params.limit, err = defaultAtoi(request.FormValue("limit"), defaultLimit); err != nil {
		return nil, errors.New("limit must be a valid integer")
	}

	if params.limit <= 0 {
		return nil, errors.New("limit must be greater than 0")
	}

	return &params, nil
}

// Our fizz buzz handler responsible for serving request.
// Only support GET requests on the root path.
func fizzBuzzHandler(responseWriter http.ResponseWriter, request *http.Request) {
	// We'll only handle GET requests
	if request.Method != "GET" {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// We'll only handle requests on the root path
	if request.URL.Path != "/" {
		responseWriter.WriteHeader(http.StatusNotFound)
	}

	// Our endpoint serves JSON content
	responseWriter.Header().Set("Content-Type", "application/json; charset=UTF-8")

	params, err := fetchAndCheckParameters(request)

	// If something went wrong with parameters send an error report to the user
	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(responseWriter).Encode(responseError{
			http.StatusBadRequest,
			err.Error(),
		})
		return
	}

	log.Println("Serving request with parameters :", params)

	fizzBuzz(params, &jsonWriterFizzBuzzCallback{responseWriter})
}