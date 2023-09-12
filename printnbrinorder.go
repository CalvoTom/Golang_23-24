package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	var tab []rune
	var min int

	if n < 0 {
		z01.PrintRune(45)
	}
	if n == 0 {
		z01.PrintRune(48)
		return
	}
	for n != 0 {
		i := 48
		for j := 1; j <= n%10; j++ {
			i += 1
		}
		for j := -1; j >= n%10; j-- {
			i += 1
		}
		n = n / 10
		tab = append(tab, rune(i))
	}
	for i := 0; i < len(tab)-1; i++ {
		min = i
		for j := i + 1; j < len(tab); j++ {
			if tab[j] < tab[min] {
				min = j
			}
		}
		if min != i {
			val := tab[i]
			tab[i] = tab[min]
			tab[min] = val
		}
	}
	for i := 0; i < len(tab); i++ {
		z01.PrintRune(tab[i])
	}
}
