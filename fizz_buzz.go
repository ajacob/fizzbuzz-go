package main

// This class contains all the computation logic to do the fizz buzz.

import (
	"strconv"
)

// Parameters to be used to do the fizz buzz computation
type Parameters struct {
	string1, string2 string
	int1, int2 int
	limit int
}

// Returns true is i is a multiple of j.
func isMultipleOf(i, j int) bool {
	return i % j == 0
}

// FizzBuzz computes the values of the sequence according to provided parameters.
// Each resulting value will be passed to callback.
func FizzBuzz(params *Parameters, callback func (value string, isLast bool)) {
	string1String2 := params.string1 + params.string2

	var nextValue string

	for i := 1; i <= params.limit; i++ {
		isMultipleOfInt1 := isMultipleOf(i, params.int1)
		isMultipleOfInt2 := isMultipleOf(i, params.int2)

		if isMultipleOfInt1 && isMultipleOfInt2 {
			nextValue = string1String2
		} else if isMultipleOfInt1 {
			nextValue = params.string1
		} else if isMultipleOfInt2 {
			nextValue = params.string2
		} else {
			nextValue = strconv.Itoa(i)
		}

		// For the callback to know if there will be more elements
		isLast := params.limit == i

		callback(nextValue, isLast)
	}
}
