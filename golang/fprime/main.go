package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 1 {
		return
	}

	first := true
	divisor := 2

	for n > 1 {
		if n%divisor == 0 {
			if !first {
				fmt.Print("*")
			}
			fmt.Print(divisor)
			n /= divisor
			first = false
		} else {
			divisor++
		}
	}

	fmt.Println()
}
