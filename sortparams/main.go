package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		z01.PrintRune(r[i])
	}
}

func main() {
	args := os.Args[1:]
	var tab []string
	var min int
	for i := 0; i < len(args); i++ {
		for j, arg := range args {
			if j >= 0 {
				tab = append(tab, arg)
			}
		}
	}
	for j := 0; j < len(tab)-1; j++ {
		min = j
		for k := j + 1; k < len(tab); k++ {
			if tab[k] < tab[min] {
				min = k
			}
		}
		if min != j {
			val := tab[j]
			tab[j] = tab[min]
			tab[min] = val
		}
	}
	for h := 0; h < len(tab); h += 3 {
		PrintStr(tab[h])
		z01.PrintRune('\n')
	}
}
