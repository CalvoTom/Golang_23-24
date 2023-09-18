package piscine

func SplitWhiteSpaces(s string) []string {
	var tabStringSeparate []string
	lastIndice := 0
	for indice, val := range s {
		if val == ' ' && s[indice-1] != ' ' {
			tabStringSeparate = append(tabStringSeparate, s[lastIndice:indice])
			if s[indice+1] != ' ' {
				lastIndice = indice + 1
			} else {
				lastIndice = indice + 2
			}
		}
		if indice == len(s)-1 {
			tabStringSeparate = append(tabStringSeparate, s[lastIndice:])
			return tabStringSeparate
		}
	}
	return tabStringSeparate
}
