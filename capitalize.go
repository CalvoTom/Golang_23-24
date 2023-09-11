package piscine

func Capitalize(s string) string {
	str_rn := []rune(s)
	if str_rn[0] >= 'a' && str_rn[0] <= 'z' {
		str_rn[0] = str_rn[0] - 32
	}
	for i := 1; i < len(str_rn); i++ {
		if str_rn[i] >= 'A' && str_rn[i] <= 'Z' {
			if str_rn[i-1] >= 'a' && str_rn[i-1] <= 'z' {
				str_rn[i] = str_rn[i] + 32
			}
			if str_rn[i-1] >= 'A' && str_rn[i-1] <= 'Z' {
				str_rn[i] = str_rn[i] + 32
			}
			if str_rn[i-1] >= '0' && str_rn[i-1] <= '9' {
				str_rn[i] = str_rn[i] + 32
			}
		}
		if str_rn[i] >= 'a' && str_rn[i] <= 'z' {
			if str_rn[i-1] >= 'a' && str_rn[i-1] <= 'z' {
				continue
			}
			if str_rn[i-1] >= 'A' && str_rn[i-1] <= 'Z' {
				continue
			}
			if str_rn[i-1] >= '0' && str_rn[i-1] <= '9' {
				continue
			}
			str_rn[i] = str_rn[i] - 32
		}
	}
	return string(str_rn)
}
