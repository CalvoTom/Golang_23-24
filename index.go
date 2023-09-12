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
			for j := 1; j < len(tab_tf); j++ {
				for k := i + 1; k < len(tab_s); k++ {
					if tab_s[k] == tab_s[j] {
						min = i
					}
				}
			}
		}
	}
	return min
}
