package piscine

func LastWord(s string) string {
	end := len(s)

	// Skip spaces at the end.
	for end > 0 && s[end-1] == ' ' {
		end--
	}

	// Only spaces.
	if end == 0 {
		return "\n"
	}

	// Find the beginning of the last word.
	start := end - 1
	for start > 0 && s[start-1] != ' ' {
		start--
	}

	return s[start:end] + "\n"
}
