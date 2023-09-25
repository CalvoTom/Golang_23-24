package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		fmt.Println(1)
	}
	if arguments[0] == "-c" && int(arguments[1][0]-48) > 0 {
		for index, ch := range arguments {
			if len(arguments) > 3 {
				if index >= 2 {
					data, err := os.ReadFile(ch)
					if err != nil {
						fmt.Println(err)
						continue
					}
					fmt.Println("==>", ch, "<==")
					for i := len(data) - 3; i < len(data); i++ {
						fmt.Printf(string(data[i]))
					}
					fmt.Println()
				}
			} else {
				if index == 2 {
					data, err := os.ReadFile(ch)
					if err != nil {
						fmt.Println(err)
						continue
					}
					for i := len(data) - 3; i < len(data); i++ {
						fmt.Printf(string(data[i]))
					}
					fmt.Println()
				}
			}
		}
	}
}
