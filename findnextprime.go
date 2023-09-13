package piscine

func FindNextPrime(nb int) int {
	if !IsPrime_test(nb) {
		nb = FindNextPrime(nb + 1)
	}
	return nb
}

func IsPrime_test(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := 2; i < nb-1; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
