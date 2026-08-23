package piscine

func RevConcatAlternate(slice1, slice2 []int) []int {
	result := []int{}

	i := len(slice1) - 1
	j := len(slice2) - 1

	// If slice2 is larger, take its extra elements first.
	if len(slice2) > len(slice1) {
		for j > i {
			result = append(result, slice2[j])
			j--
		}
	}

	// If slice1 is larger, take its extra elements first.
	if len(slice1) > len(slice2) {
		for i > j {
			result = append(result, slice1[i])
			i--
		}
	}

	// Once both slices have the same remaining length,
	// alternate starting with slice1.
	for i >= 0 && j >= 0 {
		result = append(result, slice1[i])
		result = append(result, slice2[j])
		i--
		j--
	}

	// Remaining elements, if any.
	for i >= 0 {
		result = append(result, slice1[i])
		i--
	}

	for j >= 0 {
		result = append(result, slice2[j])
		j--
	}

	return result
}
