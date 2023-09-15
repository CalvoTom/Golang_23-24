package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1]
	for i, arg := range args {
		if i >= 0 {
			z01.PrintRune(arg)
		}
	}
	z01.PrintRune('\n')
}
