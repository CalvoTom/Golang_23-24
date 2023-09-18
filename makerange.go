package piscine

func MakeRange(min, max int) []int {
	var tab []int
	if min >= max {
		return tab
	}
	tab = make([]int, min, max)
	size := max - min
	for i := 0; i < size; i++ {
		tab[i] = min
		min += 1

	}
	return tab
}
