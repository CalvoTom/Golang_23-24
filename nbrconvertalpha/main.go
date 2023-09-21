package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	var nb int
	signe := 0
	j := 0
	if s[0] == '-' {
		signe = 1
		j += 1
	}
	if s[0] == '+' {
		signe = 0
		j += 1
	}
	for i := j; i < len(s); i++ {
		if int(s[i]-48) >= 0 && int(s[i]-48) <= 9 {
			nb = (nb * 10) + int(s[i]-48)
		} else {
			return 0
		}
	}
	if signe == 1 {
		nb = -nb
	}
	return nb
}

func main() {
	var tab []int
	var nb int
	needMaj := 0
	args := os.Args[1:]
	if args[0] == "--upper" {
		needMaj = 1
	}
	if needMaj == 1 {
		for i := 1; i < len(args); i++ {
			if args[i] != " " {
				nb = Atoi(args[i]) + 96
				tab = append(tab, nb)
			} else {
				nb = Atoi(args[i]) + 96
				tab = append(tab, nb)
				tab = append(tab, 32)
				i++
			}
		}
		for j := 0; j < len(tab); j++ {
			z01.PrintRune(rune(tab[j] - 32))
		}
	} else {
		for i := 0; i < len(args); i++ {
			if args[i] != "14" {
				nb = Atoi(args[i]) + 96
				tab = append(tab, nb)
			} else {
				nb = Atoi(args[i]) + 96
				tab = append(tab, nb)
				tab = append(tab, 32)
				i += 1
			}
		}
		for j := 0; j < len(tab); j++ {
			z01.PrintRune(rune(tab[j]))
		}
	}
	z01.PrintRune('\n')
}
