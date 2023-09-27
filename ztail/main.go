package main

import (
	"fmt"
	"os"
)

func main() {
	nbrByte, tabFileName := FindArgs()
	if len(tabFileName) == 1 {
		dataFile, err := os.ReadFile(tabFileName[0])
		if err != nil {
			fmt.Print(err)
			fmt.Printf("\n")
			defer func() { os.Exit(1) }()
		}
		for i := len(dataFile) - nbrByte; i < len(dataFile); i++ {
			if string(dataFile[i]) != "\n" {
				fmt.Print(string(dataFile[i]))
			}
		}
	} else {
		for index, fileName := range tabFileName {
			dataFile, err := os.ReadFile(tabFileName[index])
			imin := len(dataFile) - nbrByte
			bouter := true
			if err != nil {
				fmt.Print(err)
				fmt.Printf("\n")
				defer func() { os.Exit(1) }()
				bouter = false
				continue
			}
			if bouter {
				fmt.Printf("\n")
				fmt.Print("==> ", fileName, " <==", "\n")
				if imin > 0 {
					for i := len(dataFile) - nbrByte; i < len(dataFile); i++ {
						if string(dataFile[i]) != "\n" {
							fmt.Print(string(dataFile[i]))
						}
					}
				} else {
					if err != nil {
						fmt.Print(err)
						fmt.Printf("\n")
						defer func() { os.Exit(1) }()
						continue
					}
					for i := 0; i < len(dataFile); i++ {
						if string(dataFile[i]) != "\n" {
							fmt.Print(string(dataFile[i]))
						}
					}
				}
			}
			fmt.Printf("\n")
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
