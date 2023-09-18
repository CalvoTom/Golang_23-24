package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	tab := make([]int, min, max)
	size := max - min
	for i := 0; i < size; i++ {
		tab[i] = min
		min += 1
	}
	return tab
}
