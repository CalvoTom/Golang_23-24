package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for g := 0; g < len(a); g++ {
		min := g
		for j := g + 1; j < len(a); j++ {
			if f(a[min], a[j]) > 0 {
				min = j
			}
		}
		tmp := a[g]
		a[g] = a[min]
		a[min] = tmp
	}
}
