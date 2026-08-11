package piscine

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	// Validate camelCase
	for i, r := range s {
		// Numbers and punctuation are not allowed.
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return s
		}

		// The word cannot end with a capital letter.
		if i == len(s)-1 && r >= 'A' && r <= 'Z' {
			return s
		}

		// Two capital letters cannot follow each other.
		if i > 0 && r >= 'A' && r <= 'Z' {
			prev := s[i-1]
			if prev >= 'A' && prev <= 'Z' {
				return s
			}
		}
	}

	// Convert camelCase to snake_case.
	result := ""

	for i, r := range s {
		if r >= 'A' && r <= 'Z' {
			if i != 0 {
				result += "_"
			}
			result += string(r)
		} else {
			result += string(r)
		}
	}

	return result
}
