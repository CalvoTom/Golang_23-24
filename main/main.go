package main

import (
	"fmt"
	"piscine"
)

func f(a, b int) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

func main() {
	a1 := []int{-838010, -826582, -426779, -88594, 237044, 457667, 705974, 901994}
	a2 := []int{0, 2, 1, 3}

	result1 := piscine.IsSorted(f, a1)
	result2 := piscine.IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
