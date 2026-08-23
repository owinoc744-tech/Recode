package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		stack := []rune{}
		valid := true

		for _, ch := range arg {
			switch ch {
			case '(', '[', '{':
				stack = append(stack, ch)

			case ')', ']', '}':
				if len(stack) == 0 {
					valid = false
					break
				}

				last := stack[len(stack)-1]

				if (ch == ')' && last != '(') ||
					(ch == ']' && last != '[') ||
					(ch == '}' && last != '{') {
					valid = false
					break
				}

				stack = stack[:len(stack)-1]
			}

			if !valid {
				break
			}
		}

		if valid && len(stack) == 0 {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}
