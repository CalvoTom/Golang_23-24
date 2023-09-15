package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	var tab_rev []rune
	for i := 0; i < len(args); i++ {
		for j, arg := range args[i] {
			if j >= 0 {
				tab_rev = append(tab_rev, arg)
			}
		}
	}
	for k := len(tab_rev) - 1; k > 0; k-- {
		z01.PrintRune(tab_rev[k])
	}
	z01.PrintRune('\n')
}
