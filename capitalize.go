package piscine

func Capitalize(s string) string {
	str_rn := []rune(s)
	for i := 1; i < len(str_rn); i++ {
		if str_rn[i] <= 90 && str_rn[i] >= 65 {
			str_rn[i] = str_rn[i] + 32
		}
		if !(str_rn[i] == '.' || str_rn[i] == '!' || str_rn[i] == '?') {
			i += 1
			continue
		}
	}
	return string(str_rn)
}
