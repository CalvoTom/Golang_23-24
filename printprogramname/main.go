package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i, arg := range args[0] {
		if i > 76 {
			z01.PrintRune(arg)
		}
	}
	z01.PrintRune('\n')
}
