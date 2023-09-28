package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var add string
	var or bool = false
	var ns string
	if len(os.Args) <= 1 {
		fmt.Println("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.")
		return
	}
	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		fmt.Println("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.")
		return
	} else if len(os.Args[1]) >= 3 {
		if os.Args[1][:3] == "-i=" {
			add += os.Args[1][3:]
		} else if os.Args[1][:9] == "--insert=" {
			add += os.Args[1][9:]
		}
		if len(os.Args[2]) >= 2 {
			if os.Args[2] == "-o" || os.Args[2] == "--order" {
				or = true
			}
		}
	}
	if len(os.Args[1]) >= 2 {
		if os.Args[1] == "-o" || os.Args[1] == "--order" {
			or = true
		}
	}
	ns += string(os.Args[len(os.Args)-1])
	ns += add
	if or {
		var tab []string
		for i := 0; i < len(ns); i++ {
			tab = append(tab, string(ns[i]))
		}
		for g := 0; g < len(tab); g++ {
			min := g
			for j := g + 1; j < len(tab); j++ {
				if tab[min] > tab[j] {
					min = j
				}
			}
			tmp := tab[g]
			tab[g] = tab[min]
			tab[min] = tmp
		}
		for g := 0; g < len(tab); g++ {
			if rune(tab[g][0]) == ' ' {
				z01.PrintRune(rune(' '))
			}
			for c := 48; c < 123; c++ {
				if rune(tab[g][0]) == rune(c) {
					z01.PrintRune(rune(c))
				}
			}
		}
		z01.PrintRune('\n')
	} else {
		fmt.Println(ns)
	}
}
