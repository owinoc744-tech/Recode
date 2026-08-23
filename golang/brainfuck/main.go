package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	code := os.Args[1]

	// 2048 bytes of memory.
	memory := make([]byte, 2048)

	// Pointer starts at the first byte.
	ptr := 0

	// Store matching brackets.
	matching := make(map[int]int)
	stack := []int{}

	for i, ch := range code {
		if ch == '[' {
			stack = append(stack, i)
		} else if ch == ']' {
			open := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			matching[open] = i
			matching[i] = open
		}
	}

	for i := 0; i < len(code); i++ {
		switch code[i] {
		case '>':
			ptr++

		case '<':
			ptr--

		case '+':
			memory[ptr]++

		case '-':
			memory[ptr]--

		case '.':
			fmt.Printf("%c", memory[ptr])

		case '[':
			if memory[ptr] == 0 {
				i = matching[i]
			}

		case ']':
			if memory[ptr] != 0 {
				i = matching[i]
			}
		}
	}
}
