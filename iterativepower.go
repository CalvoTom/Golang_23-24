package piscine

func IterativePower(nb int, power int) int {
	n := 1
	if power < 0 {
		return 0
	}
	for i := 0; i < power; i++ {
		n = n * nb
	}
	return n
}
