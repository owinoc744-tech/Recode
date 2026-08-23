package piscine

func ThirdTimeIsACharm(str string) string {
	result := ""

	for i, r := range str {
		if (i+1)%3 == 0 {
			result += string(r)
		}
	}

	return result + "\n"
}
