package piscine

func Compare(a, b string) int {
	str_rn_a := []rune(a)
	str_rn_b := []rune(b)
	minLength := min(len(str_rn_a), len(str_rn_b))
	i := 0
	for i < minLength {
		if str_rn_a[i] != str_rn_b[i] {
			break
		}
		i++
	}
	if i == minLength {
		if len(str_rn_a) > len(str_rn_b) {
			return 1
		} else if len(str_rn_a) < len(str_rn_b) {
			return -1
		} else {
			return 0
		}
	} else {
		if str_rn_a[i] > str_rn_b[i] {
			return 1
		} else {
			return -1
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
