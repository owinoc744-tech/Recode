package main

import (
	"fmt"
	"os"
)

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	word := os.Args[1]

	vowelPos := -1

	for i := 0; i < len(word); i++ {
		if isVowel(word[i]) {
			vowelPos = i
			break
		}
	}

	if vowelPos == -1 {
		fmt.Println("No vowels")
		return
	}

	if vowelPos == 0 {
		fmt.Println(word + "ay")
		return
	}

	fmt.Println(word[vowelPos:] + word[:vowelPos] + "ay")
}
