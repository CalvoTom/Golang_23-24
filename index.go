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
			min = i
			if len(tab_tf) == 1 {
				return min
			}
			if tab_s[i+1] != tab_tf[1] {
				min = -1
			}
		}
	}
	return min
}
