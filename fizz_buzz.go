package main

// This class contains all the computation logic to do the fizz buzz.
// It also contains an implementation for a callback to be used
// to write the result as JSON.

import (
	"io"
	"encoding/json"
	"strconv"
)

// Parameters to be used to do the fizz buzz computation
type parameters struct {
	string1, string2 string
	int1, int2 int
	limit int
}

// Describes callbacks to be called when computing fizz buzz.
type fizzBuzzCallback interface {
	// Start will be called before starting computation
	Start()

	// AddValue will be called with the value corresponding to
	// either string1 or string2 or string1string2
	// or the current number from the sequence (an int converted to a string)
	AddValue(value string, hasMore bool)

	// End will be called after the computation is done
	End()
}

// Computes fizz buzz according to provided parameters.
func fizzBuzz(params *parameters, callback fizzBuzzCallback) {
	callback.Start()

	string1String2 := params.string1 + params.string2

	for i := 1; i <= params.limit; i++ {
		isMultipleOfInt1 := i % params.int1 == 0;
		isMultipleOfInt2 := i % params.int2 == 0;

		hasMore := params.limit != i

		if isMultipleOfInt1 && isMultipleOfInt2 {
			callback.AddValue(string1String2, hasMore)
		} else if isMultipleOfInt1 {
			callback.AddValue(params.string1, hasMore)
		} else if isMultipleOfInt2 {
			callback.AddValue(params.string2, hasMore)
		} else {
			callback.AddValue(strconv.Itoa(i), hasMore)
		}
	}

	callback.End()
}

// Fizz buzz callback implementation that writes result of the
// computation as JSON.
type jsonWriterFizzBuzzCallback struct {
	writer io.Writer
}

// Start is called before the computation starts.
// Writes the beginning of the JSON array containing the results.
func (c *jsonWriterFizzBuzzCallback) Start() {
	c.writer.Write([]byte("["))
}

// AddValue is called when receiving one element of the result.
// Writer the element as a JSON string.
func (c *jsonWriterFizzBuzzCallback) AddValue(value string, hasMore bool) {
	if data, err := json.Marshal(value); err == nil {
		c.writer.Write(data)

		if hasMore {
			c.writer.Write([]byte(","))
		}
	} else {
		panic(err)
	}
}

// Start is called at the end of the computation.
// Writes the end of the JSON array containing the results.
func (c *jsonWriterFizzBuzzCallback) End() {
	c.writer.Write([]byte("]"))
}
