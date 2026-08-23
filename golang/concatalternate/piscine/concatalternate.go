package piscine

func ConcatAlternate(slice1, slice2 []int) []int {
	result := []int{}

	i := 0
	j := 0

	// If slice2 is larger, start with slice2.
	if len(slice2) > len(slice1) {
		for i < len(slice1) && j < len(slice2) {
			result = append(result, slice2[j])
			result = append(result, slice1[i])
			i++
			j++
		}

		for j < len(slice2) {
			result = append(result, slice2[j])
			j++
		}

		return result
	}

	// If slice1 is larger or both are equal,
	// start with slice1.
	for i < len(slice1) && j < len(slice2) {
		result = append(result, slice1[i])
		result = append(result, slice2[j])
		i++
		j++
	}

	for i < len(slice1) {
		result = append(result, slice1[i])
		i++
	}

	for j < len(slice2) {
		result = append(result, slice2[j])
		j++
	}

	return result
}
