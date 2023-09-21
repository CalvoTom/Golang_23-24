package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		help()
		return
	}
	argumentsSeparate := Split(arguments[0], "=")
	if arguments[0] == "--order" || arguments[0] == "-o" {
		order()
		z01.PrintRune('\n')
		return
	}
	if argumentsSeparate[0] == "--insert" && arguments[1] != "--order" || argumentsSeparate[0] == "-i" && arguments[1] != "--order" {
		insert(arguments[1:])
		insert(argumentsSeparate[1:])
		z01.PrintRune('\n')
		return
	}
	if arguments[0] == "--help" || arguments[0] == "-h" {
		help()
		return
	}
	if arguments[1] == "--order" && argumentsSeparate[0] == "--insert" {
		tabToOrder := insertTab(argumentsSeparate[1:])
		tabToOrder2 := orderTab()
		tabOrdering(tabToOrder2[3:], tabToOrder)
		z01.PrintRune('\n')
		return
	}
}

func insert(tab []string) {
	for i := 0; i < len(tab); i++ {
		for j, arg := range tab[i] {
			if j >= 0 {
				z01.PrintRune(arg)
			}
		}
	}
}

func insertTab(tab []string) []rune {
	var tabReturn []rune
	for i := 0; i < len(tab); i++ {
		for j, arg := range tab[i] {
			if j >= 0 {
				tabReturn = append(tabReturn, arg)
			}
		}
	}
	return tabReturn
}

func order() {
	args := os.Args[2:]
	var tab []rune
	var min int
	for j, arg := range args {
		if j >= 0 {
			for _, ch := range arg {
				tab = append(tab, ch)
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
	for h := 0; h < len(tab); h++ {
		z01.PrintRune(tab[h])
	}
}

func orderTab() []rune {
	args := os.Args[2:]
	var tab []rune
	var min int
	for j, arg := range args {
		if j >= 0 {
			for _, ch := range arg {
				tab = append(tab, ch)
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
	return tab
}

func tabOrdering(args, array []rune) {
	var tab []rune
	var min int
	for j, arg := range args {
		if j >= 0 {
			tab = append(tab, arg)
		}
	}
	for j, rn := range array {
		if j >= 0 {
			tab = append(tab, rn)
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
	for h := 0; h < len(tab); h++ {
		z01.PrintRune(tab[h])
	}
}

func help() {
	fmt.Println("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func PrintStr(s string) {
	for _, i := range s {
		z01.PrintRune(i)
	}
}

func Split(s, sep string) []string {
	var tabStringSeparate []string
	lastIndice := 0
	for indice, val := range s {
		if val == rune(sep[0]) && s[indice+len(sep)-1] == sep[len(sep)-1] {
			if lastIndice == indice {
				lastIndice = indice + len(sep)
			} else {
				tabStringSeparate = append(tabStringSeparate, s[lastIndice:indice])
				lastIndice = indice + len(sep)
			}
		}
		if indice == len(s)-1 {
			tabStringSeparate = append(tabStringSeparate, s[lastIndice:])
		}
	}
	if tabStringSeparate[len(tabStringSeparate)-1] == "" {
		return tabStringSeparate[:len(tabStringSeparate)-1]
	}
	return tabStringSeparate
}
