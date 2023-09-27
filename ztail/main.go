package main

import (
	"fmt"
	"os"
)

func ReturnStr(s string) int {
	a := 0
	for _, i := range s {
		a *= 10
		a += int(i) - 48
	}
	return a
}

func main() {
	array := os.Args[1:]
	if array[0] != "-c" {
		return
	}
	for i := 2; i <= len(os.Args)-2; i++ {
		filePath := array[i]
		content, err := os.ReadFile(filePath)
		var content2 []byte
		if err != nil {
			fmt.Print(err)
			fmt.Printf("\n")
			defer func() { os.Exit(1) }()
			continue
		}
		if i < len(os.Args) && i != 2 && err == nil {
			fmt.Printf("\n")
		}
		c := 0
		for i := range content {
			c = i
		}
		for j, k := range content {
			if j > c-ReturnStr(array[1]) {
				content2 = append(content2, k)
			}
		}
		fmt.Print("==> ", filePath, " <==", "\n")
		fmt.Printf(string(content2))
	}
}
