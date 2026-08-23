package piscine

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 || len(nbrs) > 2 {
		return nil
	}

	start := nbrs[0]
	end := len(a)

	// Convert negative start index.
	if start < 0 {
		start = len(a) + start
	}

	// If there is a second number, use it as the end.
	if len(nbrs) == 2 {
		end = nbrs[1]

		if end < 0 {
			end = len(a) + end
		}
	}

	// Invalid range.
	if start < 0 || start > len(a) || end < 0 || end > len(a) || start > end {
		return nil
	}

	return a[start:end]
}
