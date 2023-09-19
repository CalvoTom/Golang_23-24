package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) == 1 || len(a) == 0 {
		return true
	}
	testeur := f(a[0], a[1])
	if testeur < 0 {
		for i := 1; i < len(a); i++ {
			if f(a[i-1], a[i]) >= 0 {
				return false
			}
		}
	}
	if testeur > 0 {
		for i := 1; i < len(a); i++ {
			if f(a[i-1], a[i]) <= 0 {
				return false
			}
		}
	}
	if testeur == 0 {
		for i := 1; i < len(a); i++ {
			if f(a[i-1], a[i]) != 0 {
				return false
			}
		}
	}
	return true
}
