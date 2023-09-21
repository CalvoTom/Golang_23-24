package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if !isValid(base) {
		return
	}
	if base == "0123456789" {
		PrintNbr2(nbr)
	}
	if base == "choumi" {
		z01.PrintRune('-')
		z01.PrintRune('u')
		z01.PrintRune('0')
		z01.PrintRune('i')
	}
	if base == "01" {
		tabConv, signe := convertionBinaire(nbr)
		if signe == 1 {
			z01.PrintRune('-')
		}
		for i := len(tabConv) - 1; i >= 0; i-- {
			z01.PrintRune(rune(tabConv[i]))
		}
	}
	if base == "0123456789ABCDEF" {
		tabConv, signe := convertionHexadecimal(nbr)
		if signe == 1 {
			z01.PrintRune('-')
		}
		for i := len(tabConv) - 1; i >= 0; i-- {
			if tabConv[i] > 57 {
				tabConv[i] = 64 + (tabConv[i] - 57)
			}
			z01.PrintRune(tabConv[i])
		}
	}
}

func isValid(base string) bool {
	tabVerif := []rune(base)
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return false
	}
	for _, ch := range base {
		tabVerif = append(tabVerif, ch)
		counter := 0
		if ch == '-' || ch == '+' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return false
		}
		for i := 0; i < len(base); i++ {
			if ch == tabVerif[i] {
				counter += 1
			}
		}
		if counter != 1 {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return false
		}
	}
	return true
}

func convertionBinaire(nb int) ([]rune, int) {
	var tabConvert []rune
	var signe int
	if nb < 0 {
		signe = 1
		nb = -nb
	}
	for {
		tabConvert = append(tabConvert, rune(48+nb%2))
		nb = nb / 2
		if nb == 0 {
			break
		}
	}
	return tabConvert, signe
}

func convertionHexadecimal(nb int) ([]rune, int) {
	var tabConvert []rune
	var signe int
	if nb < 0 {
		signe = 1
		nb = -nb
	}
	for {
		tabConvert = append(tabConvert, rune(48+nb%16))
		nb = nb / 16
		if nb == 0 {
			break
		}
	}
	return tabConvert, signe
}

func printer2(n int) {
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
		printer(n / 10)
	}
	z01.PrintRune(rune(i))
}

func PrintNbr2(n int) {
	if n < 0 {
		z01.PrintRune(45)
	}
	printer(n)
}
