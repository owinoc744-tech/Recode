package piscine

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	// Remove spaces.
	clean := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			clean += string(str[i])
		}
	}

	if len(clean) < 5 {
		return "Invalid Input\n"
	}

	result := ""

	for i := 0; i < len(clean); i += 6 {
		if i+5 > len(clean) {
			break
		}

		if result != "" {
			result += " "
		}

		result += clean[i : i+5]
	}

	result += "\n"
	return result
}
