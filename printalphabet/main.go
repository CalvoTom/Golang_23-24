package main

import "github.com/01-edu/z01"

func main() {
	for i := 122; i > 96; i-- {
		rn := rune(i)
		z01.PrintRune(rn)
	}
	z01.PrintRune('\n')
}
