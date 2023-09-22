package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[0:]
	for i := len(args) - 1; i > 0; i-- {
		for j, arg := range args[i] {
			if j >= 0 {
				z01.PrintRune(arg)
			}
		}
		z01.PrintRune('\n')
	}
}
