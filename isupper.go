package piscine

func IsUpper(s string) bool {
	str_rn := []rune(s)
	for i := 0; i < len(str_rn); i++ {
		if str_rn[i] < 65 || str_rn[i] > 90 {
			return false
		}
	}
	return true
}
