package piscine

func IsPrintable(s string) bool {
	str_rn := []rune(s)
	for i := 0; i < len(str_rn); i++ {
		if str_rn[i] >= 126 || str_rn[i] <= 33 {
			return false
		}
	}
	return true
}
