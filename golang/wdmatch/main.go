package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	word := os.Args[1]
	source := os.Args[2]

	i := 0

	for j := 0; j < len(source) && i < len(word); j++ {
		if word[i] == source[j] {
			i++
		}
	}

	if i == len(word) {
		fmt.Println(word)
	}
}
