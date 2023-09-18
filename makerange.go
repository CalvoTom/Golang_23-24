package piscine

func MakeRange(min, max int) []int {
	var tab []int
	if min >= max {
		return tab
	}
	tab = make([]int, min, max)
	for i := 0; i < len(tab); i++ {
		tab[i] = min
		min += 1

	}
	return tab
}
