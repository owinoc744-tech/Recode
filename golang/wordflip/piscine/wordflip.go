package piscine

import "strings"

func WordFlip(str string) string {
	words := strings.Fields(str)

	if len(words) == 0 {
		if strings.TrimSpace(str) == "" && str != "" {
			return "\n"
		}
		return "Invalid Output\n"
	}

	var result strings.Builder

	for i := len(words) - 1; i >= 0; i-- {
		result.WriteString(words[i])

		if i != 0 {
			result.WriteString(" ")
		}
	}

	result.WriteString("\n")
	return result.String()
}
