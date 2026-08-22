package piscine

import "fmt"

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	result := ""

	if from <= to {
		for i := from; i <= to; i++ {
			if i > from {
				result += ", "
			}
			result += fmt.Sprintf("%02d", i)
		}
	} else {
		for i := from; i >= to; i-- {
			if i < from {
				result += ", "
			}
			result += fmt.Sprintf("%02d", i)
		}
	}

	return result + "\n"
}
