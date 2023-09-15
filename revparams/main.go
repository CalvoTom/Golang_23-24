package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	var tab_reverse []string
	for i := (len(args) - 1); i >= 0; i-- {
		tab_reverse = append(tab_reverse, args[i])
	}
	for i := 0; i < len(tab_reverse); i++ {
		for j, arg := range tab_reverse[i] {
			if j >= 0 {
				z01.PrintRune(arg)
			}
		}
		z01.PrintRune('\n')
	}

}
