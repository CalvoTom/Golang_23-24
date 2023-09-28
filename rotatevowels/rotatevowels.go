package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) <= 1 {
		z01.PrintRune('\n')
		return
	}
	var tabvoyell []string
	var tab []string
	var ns string
	tabvoyell = append(tabvoyell, "a")
	tabvoyell = append(tabvoyell, "e")
	tabvoyell = append(tabvoyell, "i")
	tabvoyell = append(tabvoyell, "o")
	tabvoyell = append(tabvoyell, "u")
	tabvoyell = append(tabvoyell, "A")
	tabvoyell = append(tabvoyell, "E")
	tabvoyell = append(tabvoyell, "I")
	tabvoyell = append(tabvoyell, "O")
	tabvoyell = append(tabvoyell, "U")
	for i := 1; i < len(os.Args); i++ {
		ns += os.Args[i]
		if i != len(os.Args)-1 {
			ns += " "
		}
	}
	for c := 0; c < len(ns); c++ {
		tab = append(tab, string(ns[c]))
	}
	var tabxyvoy []int
	for g := 0; g < len(tab); g++ {
		for j := 0; j < len(tabvoyell); j++ {
			if tab[g] == tabvoyell[j] {
				tabxyvoy = append(tabxyvoy, g)
			}
		}
	}
	comp := 1
	for h := 0; h < len(tabxyvoy)/2; h++ {
		temp := tab[tabxyvoy[len(tabxyvoy)-comp]]
		tab[tabxyvoy[len(tabxyvoy)-comp]] = tab[tabxyvoy[h]]
		tab[tabxyvoy[h]] = temp
		comp++
	}
	for p := 0; p < len(tab); p++ {
		for x := 0; x < len(tab[p]); x++ {
			z01.PrintRune(rune(tab[p][x]))
		}
	}
	z01.PrintRune('\n')
}
