package piscine

func IterativePower(nb int, power int) int {
	n := 1
	for i := 0; i < power; i++ {
		n = n * nb
	}
	return n
}
