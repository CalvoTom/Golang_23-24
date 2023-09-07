package piscine

func AlphaCount(s string) int {
	str_rn := []rune(s)
	counter := 0
	for i := 0; i < len(str_rn); i++ {
		if str_rn[i] <= 'z' && str_rn[i] >= 'a' || str_rn[i] <= 'Z' && str_rn[i] >= 'A' {
			counter += 1
		}
	}
	return counter
}
