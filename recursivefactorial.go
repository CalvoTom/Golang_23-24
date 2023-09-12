package piscine

func IterativeFactorial(nb int) int {
	if nb > 20 || nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	if nb != 1 {
		nb = nb * IterativeFactorial(nb-1)
	}
	return nb
}
