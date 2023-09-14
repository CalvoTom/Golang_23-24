package piscine

func StrRev(s string) string {
	tab := []rune(s)
	var tab_reverse []rune
	for i := (len(tab) - 1); i >= 0; i-- {
		tab_reverse = append(tab_reverse, tab[i])
	}
	return string(tab_reverse)
}
