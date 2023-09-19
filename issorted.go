package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	testeur := f(a[0], a[1])
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) != testeur {
			return false
		}
	}
	return true
}
