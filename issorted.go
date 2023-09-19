package piscine

func f(a, b int) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) != -1 {
			return false
		}
	}
	return true
}
