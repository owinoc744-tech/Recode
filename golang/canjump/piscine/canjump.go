package piscine

func CanJump(input []uint) bool {
	if len(input) == 0 {
		return false
	}

	position := 0

	for position < len(input)-1 {
		position += int(input[position])

		if position >= len(input) {
			return false
		}
	}

	return position == len(input)-1
}
