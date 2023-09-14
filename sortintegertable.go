package piscine

func SortIntegerTable(table []int) {
	var min int
	for i := 0; i < len(table)-1; i++ {
		min = i
		for j := i + 1; j < len(table); j++ {
			if table[j] < table[min] {
				min = j
			}
		}
		if min != i {
			val := table[i]
			table[i] = table[min]
			table[min] = val
		}
	}
}
