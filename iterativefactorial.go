package piscine

func IterativeFactorial(nb int) int {
	if nb > 11 {
		return 0
	}
	if nb != 1 {
		nb = nb * IterativeFactorial(nb-1)
	}
	return nb
}
