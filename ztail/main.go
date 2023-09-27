package main

import (
	"fmt"
	"os"
)

func main() {
	nbrByte, tabFileName := FindArgs()
	if len(tabFileName) == 1 {
		dataFile, err := os.ReadFile(tabFileName[0])
		check(err)
		for i := len(dataFile) - nbrByte; i < len(dataFile); i++ {
			if string(dataFile[i]) != "\n" {
				fmt.Print(string(dataFile[i]))
			}
		}
	} else {
		for index, fileName := range tabFileName {
			dataFile, err := os.ReadFile(tabFileName[index])
			imin := len(dataFile) - nbrByte

			if check(err) {
				fmt.Println()
				fmt.Println("==>", fileName, "<==")
				if imin > 0 {
					for i := len(dataFile) - nbrByte; i < len(dataFile); i++ {
						if string(dataFile[i]) != "\n" {
							fmt.Print(string(dataFile[i]))
						}
					}
					fmt.Println()
				} else {
					for i := 0; i < len(dataFile); i++ {
						if string(dataFile[i]) != "\n" {
							fmt.Print(string(dataFile[i]))
						}
					}
					fmt.Println()
					fmt.Print(fmt.Errorf("exit status 1"))
					fmt.Println()
					fmt.Println()
				}
			}
		}
	}
}

func FindArgs() (int, []string) {
	var tabFileName []string
	var nbrByte int
	arguments := os.Args[1:]
	if arguments[0] == "-c" && int(arguments[1][0]-48) > 0 {
		for _, ch := range arguments[1] {
			nbrByte = (nbrByte * 10) + int(ch-48)
		}
		for index, carac := range arguments {
			if index >= 2 {
				tabFileName = append(tabFileName, carac)
			}
		}
	}
	return nbrByte, tabFileName
}

func check(err error) bool {
	if err != nil {
		fmt.Print(err)
		fmt.Println()
		return false
	}
	return true
}
