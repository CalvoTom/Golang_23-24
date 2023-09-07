package piscine

func LastRune(s string) rune {
	r := []rune(s)
	last_indice := len(r)
	return r[last_indice-1]
}
