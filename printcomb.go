package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 48; i < 58; i++ {
		r1 := rune(i)
		for j := 48; j < 58; j++ {
			r2 := rune(j)
			for k := 48; k < 58; k++ {
				r3 := rune(k)
				if r1 < r2 && r2 < r3 {
					z01.PrintRune(r1)
					z01.PrintRune(r2)
					z01.PrintRune(r3)
					if r1 != rune(55) {
						z01.PrintRune(rune(44))
						z01.PrintRune(rune(32))
					}
				}
			}
		}
	}
	z01.PrintRune(rune('\n'))
}
