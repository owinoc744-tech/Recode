package main

import (
	"fmt"
	"math"
)

func Entropy(s string) float64 {
	length := float64(len(s))
	if length == 0 {
		return 0.0
	}

	// Step 1: Count occurrences of each byte/character
	counts := make(map[rune]float64)
	for _, char := range s {
		counts[char]++
	}

	// Step 2: Apply the Shannon Entropy formula
	var entropy float64
	for _, count := range counts {
		prob := count / length
		entropy -= prob * math.Log2(prob)
	}

	return entropy
}

func main() {
	// Test cases matching typical check tests
	fmt.Printf("%.2f\n", Entropy("hello")) // Outputs: 2.05
	fmt.Printf("%.2f\n", Entropy("abc"))   // Outputs: 1.58
	fmt.Printf("%.2f\n", Entropy(""))      // Outputs: 0.00
}
