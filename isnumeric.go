package piscine

func IsNumeric(s string) bool {
	str_rn := []rune(s)
	for i := 0; i < len(str_rn); i++ {
		if str_rn[i] < 48 || str_rn[i] > 57 {
			return false
		}
	}
	return true
}
