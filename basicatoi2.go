package piscine

func BasicAtoi2(s string) int {
	var nb int
	for i := 0; i < len(s); i++ {
		if int(s[i]-48) >= 0 && int(s[i]-48) <= 9 {
			nb = (nb * 10) + int(s[i]-48)
		} else {
			return 0
		}
	}
	return nb
}
