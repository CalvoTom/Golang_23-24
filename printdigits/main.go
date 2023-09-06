package main

import "github.com/01-edu/z01"

func main() {
	for i := 48; i < 58; i++ {
		rn := rune(i)
		z01.PrintRune(rn)
	}
	z01.PrintRune('\n')
}
