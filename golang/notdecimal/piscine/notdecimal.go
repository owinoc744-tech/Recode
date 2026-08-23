package piscine

import (
	"strconv"
	"strings"
)

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	// Check if it is a valid number.
	_, err := strconv.ParseFloat(dec, 64)
	if err != nil {
		return dec + "\n"
	}

	// No decimal point.
	if !strings.Contains(dec, ".") {
		return dec + "\n"
	}

	parts := strings.Split(dec, ".")

	// Only zero after the decimal point.
	if parts[1] == "0" {
		return dec + "\n"
	}

	// Remove the decimal point.
	result := parts[0] + parts[1]

	// Handle negative numbers such as -0.125.
	if strings.HasPrefix(dec, "-") && parts[0] == "-0" {
		result = "-" + parts[1]
	}

	// Remove leading zero when the integer part is zero.
	if parts[0] == "0" {
		result = parts[1]
	}

	return result + "\n"
}
