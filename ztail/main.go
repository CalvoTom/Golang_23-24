package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	var nb int
	if len(arguments) == 0 {
		fmt.Println(1)
	}
	if arguments[0] == "-c" && int(arguments[1][0]-48) > 0 {
		for _, carac := range arguments[1] {
			nb = (nb * 10) + int(carac-48)
		}
		for index, ch := range arguments {
			if len(arguments) > 3 {
				if index >= 2 {
					data, err := os.ReadFile(ch)
					imin := (len(data) - nb + 1)
					if err != nil {
						if index == len(arguments)-1 {
							fmt.Println(err)
							continue
						}
						fmt.Println(err)
						fmt.Printf("\n")
						continue
					}
					fmt.Println("==>", ch, "<==")
					if imin < 0 {
						for i := 0; i < len(data); i++ {
							fmt.Printf(string(data[i]))
						}
						fmt.Printf("\n")
						fmt.Printf("exit status 1")
						return
					} else {
						for i := imin; i < len(data); i++ {
							fmt.Printf(string(data[i]))
						}
						fmt.Printf("\n")
					}
				}
			} else {
				if index == 2 {
					data, err := os.ReadFile(ch)
					imin := (len(data) - nb + 1)
					if err != nil {
						if index == len(arguments)-1 {
							fmt.Println(err)
							continue
						}
						fmt.Println(err)
						fmt.Printf("\n")
						continue
					}
					fmt.Printf("\n")
					if imin < 0 {
						for i := 0; i < len(data); i++ {
							fmt.Printf(string(data[i]))
						}
						fmt.Printf("\n")
						fmt.Printf("exit status 1")
						return
					} else {
						for i := imin; i < len(data); i++ {
							fmt.Printf(string(data[i]))
						}
						fmt.Printf("\n")
					}
				}
			}
		}
	}
}
