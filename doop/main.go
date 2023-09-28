package main

import (
	"os"
)

// Regarde si le string passÃ© est un nobre
func isNumeric(s string) bool {
	var num bool = false
	var n int = 0
	array := []rune(s)
	for r := 0; r < len(array); r++ {
		if array[r] == rune(45) || array[r] == rune(43) {
			n++
		}
		for c := 48; c <= 57; c++ {
			if array[r] == rune(c) {
				n++
			}
		}
	}
	if n == len(s) {
		num = true
	}
	return num
}

// convertie string to int
func atoi(s string) int {
	ni := 0
	neg := false
	compt := 1
	valide := 0
	for i := len(s) - 1; i >= 0; i-- {
		if (rune(s[i]) >= 48 && rune(s[i]) <= 57) || rune(s[i]) == 45 || rune(s[i]) == 43 {
			if rune(s[i]) == rune(45) && i == 0 {
				neg = true
				valide++
			} else if rune(s[i]) == rune(43) && i == 0 {
				valide++
			} else if rune(s[i]) == rune(45) && i != 0 {
				return 0
			} else if rune(s[i]) == rune(43) && i != 0 {
				return 0
			}
			for c := 48; c < 58; c++ {
				if rune(s[i]) == rune(c) {
					ni += (c - 48) * compt
					compt = compt * 10
				}
			}
		} else {
			return 0
		}
	}
	if valide >= 2 {
		return 0
	}
	if neg {
		return ni * (-1)
	}
	return ni
}

func iota(n int) string {
	if n == 0 {
		return "0"
	}
	neg := false
	p := 0
	if n < 0 {
		neg = true
		p = n * (-1)
	} else {
		p = n
	}
	sn := ""
	for i := 2; i > 0; i++ {
		if p != 0 {
			m := p % 10
			if m == 0 {
				sn = "0" + sn
			} else if m == 1 {
				sn = "1" + sn
			} else if m == 2 {
				sn = "2" + sn
			} else if m == 3 {
				sn = "3" + sn
			} else if m == 4 {
				sn = "4" + sn
			} else if m == 5 {
				sn = "5" + sn
			} else if m == 6 {
				sn = "6" + sn
			} else if m == 7 {
				sn = "7" + sn
			} else if m == 8 {
				sn = "8" + sn
			} else if m == 9 {
				sn = "9" + sn
			}
			p /= 10
		} else {
			i = -99
		}
	}
	if neg {
		sn = "-" + sn
		return sn
	}
	return sn
}

func main() {
	var tab []int
	los := len(os.Args)
	if los <= 3 || los >= 5 {
		return
	} else if !isNumeric(os.Args[1]) || !isNumeric(os.Args[3]) {
		return
	}
	tab = append(tab, atoi(os.Args[1]))
	tab = append(tab, atoi(os.Args[3]))
	if tab[0] >= 922337203685477580 || tab[1] >= 922337203685477580 || tab[0] <= -922337203685477580 || tab[1] <= -922337203685477580 {
		return
	}
	if os.Args[2] == "/" {
		if tab[1] == 0 {
			os.Stdout.WriteString("No division by 0")
		} else {
			rep := tab[0] / tab[1]
			os.Stdout.WriteString(iota(rep))
		}
	} else if os.Args[2] == "*" {
		rep := tab[0] * tab[1]
		os.Stdout.WriteString(iota(rep))
	} else if os.Args[2] == "+" {
		rep := tab[0] + tab[1]
		os.Stdout.WriteString(iota(rep))
	} else if os.Args[2] == "-" {
		rep := tab[0] - tab[1]
		os.Stdout.WriteString(iota(rep))
	} else if os.Args[2] == "%" {
		if tab[1] == 0 {
			os.Stdout.WriteString("No modulo by 0")
		} else {
			rep := tab[0] % tab[1]
			os.Stdout.WriteString(iota(rep))
		}
	} else {
		return
	}
	os.Stdout.WriteString("\n")
}
