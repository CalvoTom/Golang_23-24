package piscine

func ToLower(s string) string {
	str_rn := []rune(s)
	for i := 0; i < len(str_rn); i++ {
		if str_rn[i] <= 90 && str_rn[i] >= 65 {
			str_rn[i] = str_rn[i] + 32
		}
	}
	return string(str_rn)
}
