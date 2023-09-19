package piscine

func CountIf(f func(string) bool, tab []string) int {
	counter := 0
	for _, val := range tab {
		if f(val) == true {
			counter += 1
		}
	}
	return counter
}
