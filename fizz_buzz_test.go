package main

import "testing"

// Helper to check elements of the slice with expected values.
// Using 1 based indices to facilitate human calculation.
// Element 1 match with first element of the slice (result[0])
func expectElementEqualTo(t *testing.T, result []string, element int, expected string) {
	if result[element - 1] != expected {
		t.Errorf("result[%d] is wrong, got %s instead of %s", element, result[element - 1], expected)
	}
}

// Test FizzBuzz with basic parameters.
func TestFizzBuzz(t *testing.T) {
	string1 := "fizz"
	string2 := "buzz"
	int1 := 3
	int2 := 5
	limit := 100

	params := Parameters{string1, string2, int1, int2, limit}

	var result []string

	FizzBuzz(&params, func (value string) {
		result = append(result, value)
	})

	if len(result) != limit {
		t.Errorf("Size of the result slice is %d instead of %d", len(result), limit)
	}

	expectElementEqualTo(t, result, 1, "1")
	expectElementEqualTo(t, result, 2, "2")
	expectElementEqualTo(t, result, 3, string1)
	expectElementEqualTo(t, result, 4, "4")
	expectElementEqualTo(t, result, 5, string2)
	expectElementEqualTo(t, result, 15, string1 + string2)
	expectElementEqualTo(t, result, 15, string1 + string2)
	expectElementEqualTo(t, result, 74, "74")
	expectElementEqualTo(t, result, 75, string1 + string2)
	expectElementEqualTo(t, result, 76, "76")
}

// Test FizzBuzz with different parameters.
func TestFizzBuzzDifferentParameters(t *testing.T) {
	string1 := "hello"
	string2 := "world"
	int1 := 2
	int2 := 6
	limit := 300

	params := Parameters{string1, string2, int1, int2, limit}

	var result []string

	FizzBuzz(&params, func (value string) {
		result = append(result, value)
	})

	if len(result) != limit {
		t.Errorf("Size of the result slice is %d instead of %d", len(result), limit)
	}

	expectElementEqualTo(t, result, 1, "1")
	expectElementEqualTo(t, result, 2, string1)
	expectElementEqualTo(t, result, 3, "3")
	expectElementEqualTo(t, result, 4, string1)
	expectElementEqualTo(t, result, 5, "5")
	expectElementEqualTo(t, result, 6, string1 + string2)
}

// Tests the trivial isMultipleOf helper func.
func TestIsMultipleOf(t *testing.T) {
	if !isMultipleOf(6, 3) {
		t.Errorf("6 is a multiple of 3")
	}

	if isMultipleOf(7, 3) {
		t.Errorf("7 is NOT a multiple of 3")
	}

	if !isMultipleOf(45, 5) {
		t.Errorf("45 is a multiple of 5")
	}

	if isMultipleOf(49, 5) {
		t.Errorf("49 is NOT a multiple of 5")
	}
}
