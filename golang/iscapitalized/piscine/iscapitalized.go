package piscine

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}

	startOfWord := true

	for _, r := range s {
		if r == ' ' || r == '\t' {
			startOfWord = true
			continue
		}

		if startOfWord {
			if r >= 'a' && r <= 'z' {
				return false
			}
			startOfWord = false
		}
	}

	return true
}
