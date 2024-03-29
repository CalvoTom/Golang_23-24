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
	if len(args) == 0 {
		return
	}
	if args[0] == "--upper" {
		needMaj = 1
	}
	if needMaj == 1 {
		for i := 1; i < len(args); i++ {
			nb = Atoi(args[i]) + 96
			if nb-96 == 0 || nb-96 > 26 {
				tab = append(tab, 32)
			} else {
				tab = append(tab, nb-32)
			}
		}
	} else {
		for i := 0; i < len(args); i++ {
			nb = Atoi(args[i]) + 96
			if nb-96 == 0 || nb-96 > 26 {
				tab = append(tab, 32)
			} else {
				tab = append(tab, nb)
			}
		}
	}
	for j := 0; j < len(tab); j++ {
		z01.PrintRune(rune(tab[j]))
	}
	z01.PrintRune('\n')
}
