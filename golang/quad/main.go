package piscine

import "github.com/01-edu/z01"

func printLine(x int, left, middle, right rune) {
	for i := 0; i < x; i++ {
		if i == 0 {
			z01.PrintRune(left)
		} else if i == x-1 {
			z01.PrintRune(right)
		} else {
			z01.PrintRune(middle)
		}
	}
	z01.PrintRune('\n')
}

func printQuad(x, y int, topLeft, topMiddle, topRight, side, bottomLeft, bottomMiddle, bottomRight rune) {
	if x <= 0 || y <= 0 {
		return
	}

	if y == 1 {
		printLine(x, topLeft, topMiddle, topRight)
		return
	}

	printLine(x, topLeft, topMiddle, topRight)

	for row := 1; row < y-1; row++ {
		printLine(x, side, ' ', side)
	}

	printLine(x, bottomLeft, bottomMiddle, bottomRight)
}

func QuadA(x, y int) {
	printQuad(x, y, 'o', '-', 'o', '|', 'o', '-', 'o')
}

func QuadB(x, y int) {
	printQuad(x, y, '/', '*', '\\', '*', '\\', '*', '/')
}

func QuadC(x, y int) {
	printQuad(x, y, 'A', 'B', 'A', 'B', 'C', 'B', 'C')
}

func QuadD(x, y int) {
	printQuad(x, y, 'A', 'B', 'C', 'B', 'A', 'B', 'C')
}

func QuadE(x, y int) {
	printQuad(x, y, 'A', 'B', 'C', 'B', 'C', 'B', 'A')
}
