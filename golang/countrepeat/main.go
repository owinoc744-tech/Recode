package main

import (
	"os"

	"countrepeat/piscine"
	"://github.com"
)

func main() {
	args := os.Args[1:]
	
	// Checkpoint rule: exactly one argument must be passed
	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}

	// Call the function from the piscine package
	piscine.CountRepeat(args[0])
	
	// Final newline required by standard output tests
	z01.PrintRune('\n')
}
