package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := 48; i < 58; i++ {
		r1 := rune(i)
		for j := 48; j < 58; j++ {
			r2 := rune(j)
			for k := 48; k < 58; k++ {
				r3 := rune(k)
				for l := 48; l < 58; l++ {
					r4 := rune(l)
					if r1 < r3 {
						z01.PrintRune(r1)
						z01.PrintRune(r2)
						z01.PrintRune(rune(32))
						z01.PrintRune(r3)
						z01.PrintRune(r4)
						if r1 != rune(57) || r2 != rune(56) {
							z01.PrintRune(rune(32))
							z01.PrintRune(rune(44))
						}
					}
					if r1 == r3 && r2 < r4 {
						z01.PrintRune(r1)
						z01.PrintRune(r2)
						z01.PrintRune(rune(32))
						z01.PrintRune(r3)
						z01.PrintRune(r4)
						if r1 != rune(57) || r2 != rune(56) {
							z01.PrintRune(rune(32))
							z01.PrintRune(rune(44))
						}
					}
				}
			}
		}
	}
	z01.PrintRune(rune('\n'))
}
