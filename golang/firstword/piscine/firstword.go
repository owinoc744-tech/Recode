package piscine

func FirstWord(s string) string {
	if s == "" {
		return "\n"
	}

	end := 0

	for end < len(s) && s[end] != ' ' {
		end++
	}

	return s[:end] + "\n"
}
