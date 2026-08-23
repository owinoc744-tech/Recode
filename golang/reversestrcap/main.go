package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func process(s string) string {
	runes := []rune(strings.ToLower(s))

	inWord := false
	last := -1

	for i, r := range runes {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			inWord = true
			last = i
		} else if inWord {
			runes[last] = unicode.ToUpper(runes[last])
			inWord = false
		}
	}

	if inWord {
		runes[last] = unicode.ToUpper(runes[last])
	}

	return string(runes)
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(process(os.Args[i]))
	}
}
