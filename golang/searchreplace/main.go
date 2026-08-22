package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	str := os.Args[1]
	old := os.Args[2]
	new := os.Args[3]

	if len(old) != 1 || len(new) != 1 {
		return
	}

	found := false
	result := ""

	for _, r := range str {
		if string(r) == old {
			result += new
			found = true
		} else {
			result += string(r)
		}
	}

	fmt.Println(result)

	_ = found
}
