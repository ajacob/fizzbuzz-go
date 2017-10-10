package main

import (
	"testing"
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"net/http/httptest"
)

// Tests fetchAndCheckParameters without overriding defaults.
func TestFetchAndCheckParameters(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	params, err := fetchAndCheckParameters(request)

	if err != nil {
		t.Fatal(err)
	}

	if params.string1 != defaultString1 {
		t.Errorf("expected string1 to be %s instead of %s", defaultString1, params.string1)
	}

	if params.string2 != defaultString2 {
		t.Errorf("expected string2 to be %s instead of %s", defaultString2, params.string2)
	}

	if params.int1 != defaultInt1 {
		t.Errorf("expected int1 to be %d instead of %d", defaultInt1, params.int1)
	}

	if params.int2 != defaultInt2 {
		t.Errorf("expected int2 to be %d instead of %d", defaultInt2, params.int2)
	}

	if params.limit != defaultLimit {
		t.Errorf("expected limit to be %d instead of %d", defaultLimit, params.limit)
	}
}

// Tests fetchAndCheckParameters with custom parameters to override defaults.
func TestFetchAndCheckParametersWithCustomParams(t *testing.T) {
	expectedString1 := "hello"
	expectedString2 := "world"
	expectedInt1 := 2
	expectedInt2 := 6
	expectedLimit := 50

	values := url.Values{}
	values.Set("string1", expectedString1)
	values.Set("string2", expectedString2)
	values.Set("int1", strconv.Itoa(expectedInt1))
	values.Set("int2", strconv.Itoa(expectedInt2))
	values.Set("limit", strconv.Itoa(expectedLimit))

	request, err := http.NewRequest("GET", "/?" + values.Encode(), nil)

	if err != nil {
		t.Fatal(err)
	}

	params, err := fetchAndCheckParameters(request)

	if err != nil {
		t.Fatal(err)
	}

	if params.string1 != expectedString1 {
		t.Errorf("expected string1 to be %s instead of %s", expectedString1, params.string1)
	}

	if params.string2 != expectedString2 {
		t.Errorf("expected string2 to be %s instead of %s", expectedString2, params.string2)
	}

	if params.int1 != expectedInt1 {
		t.Errorf("expected int1 to be %d instead of %d", expectedInt1, params.int1)
	}

	if params.int2 != expectedInt2 {
		t.Errorf("expected int2 to be %d instead of %d", expectedInt2, params.int2)
	}

	if params.limit != expectedLimit {
		t.Errorf("expected limit to be %d instead of %d", expectedLimit, params.limit)
	}
}

// Helper to test fetchAndCheckParameters.
// Expect the method to return an error and check with expectedMessage.
func expectFetchAndCheckParametersFailure(t *testing.T, values url.Values, expectedMessage string) {
	request, err := http.NewRequest("GET", "/?" + values.Encode(), nil)

	if err != nil {
		t.Fatal(err)
	}

	_, err = fetchAndCheckParameters(request)

	if err == nil {
		t.Errorf("Expected fetchAndCheckParameters to return an error")
	}

	if err.Error() != expectedMessage {
		t.Errorf("Expected error message to be %s instead of %s", expectedMessage, err.Error())
	}
}

func TestFetchAndCheckParametersInvalidInt1(t *testing.T) {
	values := url.Values{}
	values.Set("int1", "invalid")

	expectFetchAndCheckParametersFailure(t, values, errorInt1InvalidInteger)
}

func TestFetchAndCheckParametersInvalidInt2(t *testing.T) {
	values := url.Values{}
	values.Set("int2", "invalid")

	expectFetchAndCheckParametersFailure(t, values, errorInt2InvalidInteger)
}

func TestFetchAndCheckParametersInt1Zero(t *testing.T) {
	values := url.Values{}
	values.Set("int1", "0")

	expectFetchAndCheckParametersFailure(t, values, errorInt1OrInt2Zero)
}

func TestFetchAndCheckParametersInt2Zero(t *testing.T) {
	values := url.Values{}
	values.Set("int2", "0")

	expectFetchAndCheckParametersFailure(t, values, errorInt1OrInt2Zero)
}

func TestFetchAndCheckParametersInvalidLimit(t *testing.T) {
	values := url.Values{}
	values.Set("limit", "invalid")

	expectFetchAndCheckParametersFailure(t, values, errorLimitInvalidInteger)
}

func TestFetchAndCheckParametersLimitLowerOrEqualToZero(t *testing.T) {
	values := url.Values{}
	values.Set("limit", "0")

	expectFetchAndCheckParametersFailure(t, values, errorLimitLowerOrEqualToZero)
}

// Helper to test fizzBuzzHandler.
// Returns a ResponseRecorder corresponding to the reply of the handler.
func getResponseFromFizzBuzzHandler(t *testing.T, method string, url string) *httptest.ResponseRecorder {
	request, err := http.NewRequest(method, url, nil)

	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(fizzBuzzHandler)

	handler.ServeHTTP(responseRecorder, request)

	return responseRecorder
}

// Test fizzBuzzHandler in normal condition.
func TestFizzBuzzHandler(t *testing.T) {
	responseRecorder := getResponseFromFizzBuzzHandler(t, "GET", "/")

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler responsed with status %d instead of %d", status, http.StatusOK)
	}

	if contentType := responseRecorder.Header().Get(contentTypeHeader); contentType != contentTypeJSONUtf8 {
		t.Errorf("handler responsed with Content-Type %s instead of %s", contentType, contentTypeJSONUtf8)
	}
}

// Test fizzBuzzHandler with invalid method.
func TestFizzBuzzHandlerInvalidMethod(t *testing.T) {
	responseRecorder := getResponseFromFizzBuzzHandler(t, "POST", "/")

	if status := responseRecorder.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler responsed with status %d instead of %d", status, http.StatusOK)
	}
}

// Test fizzBuzzHandler with invalid method.
func TestFizzBuzzHandlerInvalidPath(t *testing.T) {
	responseRecorder := getResponseFromFizzBuzzHandler(t, "GET", "/invalid")

	if status := responseRecorder.Code; status != http.StatusNotFound {
		t.Errorf("handler responsed with status %d instead of %d", status, http.StatusOK)
	}
}

// Test fizzBuzzHandler with incorrect parameters.
func TestFizzBuzzHandlerInvalidParameters(t *testing.T) {
	responseRecorder := getResponseFromFizzBuzzHandler(t, "GET", "/?limit=0")

	if status := responseRecorder.Code; status != http.StatusBadRequest {
		t.Errorf("handler responsed with status %d instead of %d", status, http.StatusBadRequest)
	}

	var errorJSON map[string]interface{}

	json.Unmarshal(responseRecorder.Body.Bytes(), &errorJSON)

	if int(errorJSON["status"].(float64)) != http.StatusBadRequest {
		t.Errorf("Expected status to be %d instead of %d", http.StatusBadRequest, errorJSON["status"])
	}

	if errorJSON["message"] == "" {
		t.Error("Expected message not to be empty")
	}
}

// Tests the jsonFizzBuzz func ensuring it outputs valid JSON.
func TestJsonFizzBuzz(t *testing.T) {
	string1 := "fizz"
	string2 := "buzz"
	int1 := 3
	int2 := 5
	limit := 10

	params := Parameters{string1, string2, int1, int2, limit}

	jsonBuffer := new(bytes.Buffer)

	jsonFizzBuzz(&params, jsonBuffer)

	var unmarshalledResult []string

	// check for json validity
	if err := json.Unmarshal(jsonBuffer.Bytes(), &unmarshalledResult); err != nil {
		t.Fatal(err)
	}

	// We don't care about the actual result here, but at least check we have
	// the expected number of elements in the resulting slice
	if len(unmarshalledResult) != limit {
		t.Errorf("Expected a slice of %d elements, have %d", limit, len(unmarshalledResult))
	}
}
