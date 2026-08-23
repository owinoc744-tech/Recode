package piscine

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	chars := make(map[rune]int)

	for _, r := range str1 {
		chars[r]++
	}

	for _, r := range str2 {
		chars[r]++
	}

	count := 0

	for _, n := range chars {
		if n == 1 {
			count++
		}
	}

	return count
}
