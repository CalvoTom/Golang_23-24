package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		for j, arg := range args[i] {
			if j >= 0 {
				z01.PrintRune(arg)
			}
		}
		z01.PrintRune('\n')
	}
}
