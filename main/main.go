package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := "HelloHAhowHAareHAyou?HA"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
}
