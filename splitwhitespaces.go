package piscine

func SplitWhiteSpaces(s string) []string {
	var tabStringSeparate []string
	lastIndice := 0
	for indice, val := range s {
		if val == ' ' {
			tabStringSeparate = append(tabStringSeparate, s[lastIndice:indice])
			lastIndice = indice
		}
		if indice == len(s)-1 {
			tabStringSeparate = append(tabStringSeparate, s[lastIndice:])
		}
	}
	return tabStringSeparate
}
