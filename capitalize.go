package piscine

func Capitalize(s string) string {
	str_rn := []rune(s)
	need_maj := 1
	if str_rn[0] <= 90 && str_rn[0] >= 65 {
		need_maj = 0
	}
	for i := 1; i < len(str_rn); i++ {
		if str_rn[i] <= 90 && str_rn[i] >= 65 && need_maj == 0 {
			str_rn[i] = str_rn[i] + 32
		}
		if str_rn[i] <= 122 && str_rn[i] >= 97 && need_maj == 1 {
			str_rn[i] = str_rn[i] - 32
			need_maj = 0
		}
		if str_rn[i] <= 90 && str_rn[i] >= 65 && need_maj == 1 {
			need_maj = 0
		}
		if str_rn[i] == '.' || str_rn[i] == '!' || str_rn[i] == '?' || str_rn[i] == '}' || str_rn[i] == '>' || str_rn[i] == '[' || str_rn[i] == ']' || str_rn[i] == '^' {
			need_maj = 1
		}
	}
	return string(str_rn)
}
