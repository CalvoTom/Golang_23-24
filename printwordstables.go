package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(a []string) {
	var tabRune []rune
	for _, values := range a {
		for i := 0; i < len(values); i++ {
			tabRune = append(tabRune, rune(values[i]))
		}
		tabRune = append(tabRune, 32)
	}
	for j := 0; j < len(tabRune); j++ {
		if j == len(tabRune)-1 {
			z01.PrintRune('\n')
			break
		}
		if tabRune[j] == 32 {
			z01.PrintRune('\n')
			j += 1
		}
		z01.PrintRune(tabRune[j])
	}
}
