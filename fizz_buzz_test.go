package main

import "testing"

type StubFizzBuzzCallback struct {
	values []string
}

func (c *StubFizzBuzzCallback) Start() {

}

func (c *StubFizzBuzzCallback) AddValue(value string, hasMore bool) {
	c.values = append(c.values, value)
}

func (c *StubFizzBuzzCallback) End() {

}

func expectElementEqualTo(t *testing.T, result []string, element int, expected string) {
	if result[element] != expected {
		t.Errorf("result[%d] is wrong, got %s instead of %s", element, result[element], expected)
	}
}

func TestFizzBuzz(t *testing.T) {
	string1 := "fizz"
	string2 := "buzz"
	int1 := 3
	int2 := 5
	limit := 100

	params := parameters{string1, string2, int1, int2, limit}

	stubCallback := StubFizzBuzzCallback{}

	fizzBuzz(&params, &stubCallback)

	if len(stubCallback.values) != limit {
		t.Errorf("Size of the result slice is %d instead of %d", len(stubCallback.values), limit)
	}

	expectElementEqualTo(t, stubCallback.values, 0, "1")
	expectElementEqualTo(t, stubCallback.values, 1, "2")
	expectElementEqualTo(t, stubCallback.values, 2, string1)
	expectElementEqualTo(t, stubCallback.values, 3, "4")
	expectElementEqualTo(t, stubCallback.values, 4, string2)
	expectElementEqualTo(t, stubCallback.values, 14, string1 + string2)
}

func TestFizzBuzzDifferentParameters(t *testing.T) {
	string1 := "hello"
	string2 := "world"
	int1 := 2
	int2 := 6
	limit := 300

	params := parameters{string1, string2, int1, int2, limit}

	stubCallback := StubFizzBuzzCallback{}

	fizzBuzz(&params, &stubCallback)

	if len(stubCallback.values) != limit {
		t.Errorf("Size of the result slice is %d instead of %d", len(stubCallback.values), limit)
	}

	expectElementEqualTo(t, stubCallback.values, 0, "1")
	expectElementEqualTo(t, stubCallback.values, 1, string1)
	expectElementEqualTo(t, stubCallback.values, 2, "3")
	expectElementEqualTo(t, stubCallback.values, 3, string1)
	expectElementEqualTo(t, stubCallback.values, 4, "5")
	expectElementEqualTo(t, stubCallback.values, 5, string1 + string2)
}
