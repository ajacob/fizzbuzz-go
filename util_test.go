package main

import "testing"

func TestDefaultStringFallback(t *testing.T) {
	expected := "two"

	result := defaultString("", expected)

	if result != expected {
		t.Errorf("defaultString should use fallback, got %s instead of %s", result, expected)
	}
}

func TestDefaultStringNoFallback(t *testing.T) {
	expected := "one"

	result := defaultString(expected, "two")

	if result != expected {
		t.Errorf("defaultString should not use fallback, got %s instead of %s", result, expected)
	}
}

func TestDefaultAtoiFallback(t *testing.T) {
	expected := 10

	if result, err := defaultAtoi("", 10); err != nil {
		t.Fatal(err)
	} else if result != expected {
		 t.Errorf("defaultAtoi should use fallback, got %d instead of %d", result, expected)
	}
}

func TestDefaultAtoiNoFallback(t *testing.T) {
	expected := 6

	if result, err := defaultAtoi("6", 10); err != nil {
		t.Fatal(err)
	} else if result != expected {
		t.Errorf("defaultAtoi should not use fallback, got %d instead of %d", result, expected)
	}
}
