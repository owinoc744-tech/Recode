package main

import (
	"fmt"
	"os"
)

func printOptions() {
	fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
}

func printBits(options int) {
	for i := 31; i >= 0; i-- {
		fmt.Printf("%d", (options>>i)&1)

		if i%8 == 0 && i != 0 {
			fmt.Print(" ")
		}
	}

	fmt.Println()
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printOptions()
		return
	}

	options := 0

	for _, arg := range args {
		if len(arg) < 2 || arg[0] != '-' {
			fmt.Println("Invalid Option")
			return
		}

		for _, c := range arg[1:] {
			if c == 'h' {
				printOptions()
				return
			}

			if c < 'a' || c > 'z' {
				fmt.Println("Invalid Option")
				return
			}

			bit := int(c - 'a')
			options |= 1 << bit
		}
	}

	printBits(options)
}
