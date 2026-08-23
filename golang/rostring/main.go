package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	words := strings.Fields(os.Args[1])

	if len(words) == 0 {
		fmt.Println()
		return
	}

	for i := 1; i < len(words); i++ {
		if i > 1 {
			fmt.Print(" ")
		}
		fmt.Print(words[i])
	}

	fmt.Print(" ", words[0])
	fmt.Println()
}
