package piscine

func Index(s string, toFind string) int {
	tab_s := []rune(s)
	tab_tf := []rune(toFind)
	min := -1
	if len(tab_tf) == 0 {
		return 0
	}
	for i := 0; i < len(tab_s); i++ {
		if tab_s[i] == tab_tf[0] {
			if len(tab_tf) == 1 {
				min = i
				return min
			}
			j := 1
			for k := i + 1; j < len(tab_tf); k++ {
				if tab_s[k] != tab_tf[j] {
					min = -1
					break
				} else if j == len(tab_tf)-1 {
					return i
				}
				j++
			}
		}
	}
	return min
}
