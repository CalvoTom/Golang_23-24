package piscine

func TrimAtoi(s string) int {
	var nb int
	signe := 0
	av := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j <= 9; j++ {
			if int(s[i]-48) == j {
				av = 1
				nb = (nb * 10) + int(s[i]-48)
			}
		}
		if s[i] == '-' && av == 0 {
			signe = 1
		}
	}
	if signe == 1 {
		nb = -nb
	}
	return nb
}
