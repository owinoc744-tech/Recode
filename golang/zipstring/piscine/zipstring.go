package piscine

import "strconv"

func ZipString(s string) string {
	if s == "" {
		return ""
	}

	result := ""

	for i := 0; i < len(s); {
		count := 1

		for i+count < len(s) && s[i] == s[i+count] {
			count++
		}

		result += strconv.Itoa(count) + string(s[i])
		i += count
	}

	return result
}
