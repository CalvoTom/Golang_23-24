package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[0]
	for i, arg := range args[1:] {
		if i > 75 {
			z01.PrintRune(arg)
		}
	}
	z01.PrintRune('\n')
}
