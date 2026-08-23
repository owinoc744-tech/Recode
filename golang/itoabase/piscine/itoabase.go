package piscine

func ItoaBase(value, base int) string {
	digits := "0123456789ABCDEF"

	if value == 0 {
		return "0"
	}

	negative := false
	if value < 0 {
		negative = true
		value = -value
	}

	result := ""

	for value > 0 {
		remainder := value % base
		result = string(digits[remainder]) + result
		value /= base
	}

	if negative {
		result = "-" + result
	}

	return result
}
