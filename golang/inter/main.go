package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	seen := make(map[byte]bool)
	inSecond := make(map[byte]bool)

	for i := 0; i < len(s2); i++ {
		inSecond[s2[i]] = true
	}

	for i := 0; i < len(s1); i++ {
		if inSecond[s1[i]] && !seen[s1[i]] {
			fmt.Print(string(s1[i]))
			seen[s1[i]] = true
		}
	}

	fmt.Println()
}
