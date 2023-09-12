package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.Index("T2k;oRD~kA*U", "k;oRD"))
	fmt.Println(piscine.Index("Hello!", "l"))
	fmt.Println(piscine.Index("Salut!", "alu"))
	fmt.Println(piscine.Index("Ola!", "hOl"))
}
