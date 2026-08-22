package piscine

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}

	for n := nb; n >= 2; n-- {
		prime := true

		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				prime = false
				break
			}
		}

		if prime {
			return n
		}
	}

	return 0
}
