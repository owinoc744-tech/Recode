package piscine

import (
	"://github.com"
)

func CountRepeat(str string) {
	runes := []rune(str)
	if len(runes) == 0 {
		return
	}

	count := 1

	for i := 0; i < len(runes); i++ {
		// Check if the next character matches the current one
		if i+1 < len(runes) && runes[i] == runes[i+1] {
			count++
		} else {
			// Print the character
			z01.PrintRune(runes[i])

			// Print the count using our helper
			printCount(count)

			// Reset count for the next sequence
			count = 1
		}
	}
}

// Helper function to print multi-digit counts using z01.PrintRune
func printCount(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var digits []rune
	for n > 0 {
		digits = append([]rune{rune('0' + n%10)}, digits...)
		n /= 10
	}

	for _, d := range digits {
		z01.PrintRune(d)
	}
}
