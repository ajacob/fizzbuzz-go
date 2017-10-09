package main

import "strconv"

// Returns fallback if value is empty, value otherwise.
func defaultString(value, fallback string) string {
	if value == "" {
		return fallback
	}

	return value
}

// Returns fallback if value is empty, the result of strconv.Atoi otherwise.
func defaultAtoi(value string, fallback int) (int, error) {
	if value == "" {
		return fallback, nil
	}

	return strconv.Atoi(value)
}
