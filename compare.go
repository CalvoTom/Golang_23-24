package piscine

func Compare(a, b string) int {
	tab_a := []rune(a)
	tab_b := []rune(b)
	longeur_tab_A := len(tab_a)
	longeur_tab_B := len(tab_b)

	if tab_a[0] != tab_b[0] {
		return -1
	}
	if longeur_tab_A == longeur_tab_B {
		return 0
	} else {
		return 1
	}
}
