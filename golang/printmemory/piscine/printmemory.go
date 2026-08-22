package piscine

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdef"

	for i, b := range arr {
		z01.PrintRune(rune(hex[b/16]))
		z01.PrintRune(rune(hex[b%16]))

		if i%4 == 3 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}
	}

	for i := len(arr); i < 12; i++ {
		z01.PrintRune(' ')
		if i%4 == 3 {
			z01.PrintRune('\n')
		}
	}

	for _, b := range arr {
		if b >= 32 && b <= 126 {
			z01.PrintRune(rune(b))
		} else {
			z01.PrintRune('.')
		}
	}

	z01.PrintRune('\n')
}
