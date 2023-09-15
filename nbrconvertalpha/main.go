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
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		nb = Atoi(args[i])
		tab = append(tab, nb)
	}
	for j := 0; j < len(tab); j++ {
		if tab[j] > 26 {
			z01.PrintRune(rune(tab[j]))
		} else {
			rn := 96 + tab[j]
			z01.PrintRune(rune(rn))
		}
	}
}
