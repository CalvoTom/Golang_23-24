package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	var m int
	fin := false
	var bn int = len(base)
	sn := ""
	fn := ""
	if nbr < (-9223372036854775807) {
		PrintNbr(nbr * (-1))
	} else {
		if nbr <= 9970467 && nbr >= -9223372036854775807 {
			if nbr >= 0 {
				m = nbr
			} else {
				m = nbr * (-1)
			}
			if len(base) <= 1 {
				z01.PrintRune(rune(78))
				z01.PrintRune(rune(86))
			} else {
				for i := 0; i < len(base); i++ {
					if base[i] == 43 || base[i] == 45 {
						fin = true
					} else {
						for j := i + 1; j < len(base); j++ {
							if base[i] == base[j] {
								fin = true
							}
						}
					}
				}
				if !fin {
					for e := 0; e == 0; e = e + 0 {
						if m/bn >= 1 {
							sn += string(base[m%bn])
							m = m / bn
						} else {
							sn += string(base[m%bn])
							e = e + 1
						}
					}
					for d := len(sn) - 1; d >= 0; d-- {
						fn += string(sn[d])
					}
					if nbr < 0 {
						z01.PrintRune(rune(45))
					}
					PrintStr(fn)
				} else {
					z01.PrintRune(rune(78))
					z01.PrintRune(rune(86))
				}
			}
		} else {
			z01.PrintRune(rune(78))
			z01.PrintRune(rune(86))
		}
	}
}
