package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println()
		return
	}

	seen := make(map[byte]bool)

	for i := 1; i <= 2; i++ {
		for j := 0; j < len(os.Args[i]); j++ {
			char := os.Args[i][j]

			if !seen[char] {
				fmt.Print(string(char))
				seen[char] = true
			}
		}
	}

	fmt.Println()
}
