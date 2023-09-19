package piscine

func SplitWhiteSpaces(s string) []string {
	var tabStringSeparate []string
	lastIndice := 0
	for indice, val := range s {
		if val == ' ' {
			if lastIndice == indice {
				lastIndice = indice + 1
			} else {
				tabStringSeparate = append(tabStringSeparate, s[lastIndice:indice])
				lastIndice = indice + 1
			}
		}
		if indice == len(s)-1 {
			if lastIndice != indice {
				return tabStringSeparate
			} else {
				tabStringSeparate = append(tabStringSeparate, s[lastIndice:])
			}
		}
	}
	return tabStringSeparate
}
