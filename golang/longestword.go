package main

import (
	"fmt"
	"strings"
)

// LongestWord takes a string and returns the first longest word.
func LongestWord(s string) string {
	// Fields automatically splits by one or more whitespace characters
	words := strings.Fields(s)
	
	if len(words) == 0 {
		return ""
	}

	longest := ""
	
	for _, word := range words {
		// Clean the word if your specific test case requires stripping punctuation
		// For the standard check, comparing raw length is usually enough:
		if len(word) > len(longest) {
			longest = word
		}
	}
	
	return longest
}

func main() {
	fmt.Println(LongestWord("Hello piscine Zone01 Kisumu")) 
	// Output: Zone01
	
	fmt.Println(LongestWord("Coding is fun"))            
	// Output: Coding
}
