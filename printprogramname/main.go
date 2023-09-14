package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for _, arg := range args[0] {
		z01.PrintRune(arg)
	}
	z01.PrintRune('\n')
}
