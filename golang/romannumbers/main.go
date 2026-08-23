package main

import (
	"fmt"
	"os"
	"strconv"
)

type romanPart struct {
	value int
	roman string
	calc  string
}

func main() {
	if len(os.Args) != 2 {
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil || n <= 0 || n >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	parts := []romanPart{
		{1000, "M", "M"},
		{900, "CM", "(M-C)"},
		{500, "D", "D"},
		{400, "CD", "(D-C)"},
		{100, "C", "C"},
		{90, "XC", "(C-X)"},
		{50, "L", "L"},
		{40, "XL", "(L-X)"},
		{10, "X", "X"},
		{9, "IX", "(X-I)"},
		{5, "V", "V"},
		{4, "IV", "(V-I)"},
		{1, "I", "I"},
	}

	roman := ""
	calculation := ""

	for _, part := range parts {
		for n >= part.value {
			n -= part.value
			roman += part.roman

			if calculation != "" {
				calculation += "+"
			}
			calculation += part.calc
		}
	}

	fmt.Println(calculation)
	fmt.Println(roman)
}
