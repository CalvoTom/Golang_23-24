package piscine

func Any(f func(string) bool, a []string) bool {
	for _, val := range a {
		if f(val) == true {
			return true
		}
	}
	return false
}
