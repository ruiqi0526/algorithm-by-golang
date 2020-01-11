package main

import (
	"fmt"
	"math"
)

func Max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func MaxCover(x []int, n int) int {
	length := 1
	bound := n - 1

	for i := n - 2; i >= 0; i-- {
		for ; x[i] > int(math.Abs(float64(x[bound]))); bound-- {}
		length = Max(length, bound - i + 1)
	}

	return length
}

func main() {
	var (
		x = []int{1, 6, 2, 1, -2, 3, 5, 2, -4, 3}
		n = len(x) / 4
	)

	fmt.Println("Maximum Cover Interval of an Array")
	fmt.Println("==================================")
	fmt.Println("Given Array :")
	fmt.Printf("%v\n", x)
	fmt.Printf("Maximum Interval Length = %d\n", MaxCover(x, n))
}