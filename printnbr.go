package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune(45)
		n = -n
	}
	i := 48
	if n == 0 {
		z01.PrintRune(48)
		return
	}
	for j := 1; j <= n%10; j++ {
		i += 1
	}
	for j := -1; j >= n%10; j-- {
		i += 1
	}
	if n/10 != 0 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(rune(i))
}
