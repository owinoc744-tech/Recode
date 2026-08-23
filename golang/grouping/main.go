package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 3 || os.Args[2] == "" {
		return
	}

	expression := os.Args[1]
	text := os.Args[2]

	// The expression must be a valid regular expression.
	re, err := regexp.Compile(expression)
	if err != nil {
		return
	}

	words := strings.Fields(text)
	count := 0

	for _, word := range words {
		matches := re.FindAllStringIndex(word, -1)

		for range matches {
			count++
			fmt.Printf("%d: %s\n", count, word)
		}
	}
}
