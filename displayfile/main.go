package main

import (
	"fmt"
	"os"
)

func main() {
	var programmeName []rune
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		for j, arg := range args[i] {
			if j >= 0 {
				programmeName = append(programmeName, arg)
			}
		}
	}
	dat, err := os.ReadFile(string(programmeName))
	check(err, args)
	fmt.Print(string(dat))
}

func check(err error, args []string) {
	fmt.Println(err)
	if err != nil {
		if len(args) == 0 {
			fmt.Println("File name missing")
		} else {
			fmt.Println("Too many arguments")
		}
	}
}
