package piscine

func SaveAndMiss(arg string, num int) string {
	if num <= 0 {
		return arg
	}

	result := ""

	for i := 0; i < len(arg); i += num * 2 {
		end := i + num

		if end > len(arg) {
			end = len(arg)
		}

		result += arg[i:end]
	}

	return result
}
