package main

import (
	"fmt"
)

func Max(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

func Min(x, y int) int {
	if  x >= y {
		return x
	} else {
		return y
	}
}

func MaxProd(x []int, n int) int {
	var (
		maxEndingHere = 1
		minEndingHere = 1
		maxSoFar = 1
		temp int
	)

	for i := 0; i < n; i++ {
		if x[i] > 0 {
			maxEndingHere *= x[i]
			minEndingHere = Min(minEndingHere * x[i], 1)
		} else if x[i] == 0 {
			maxEndingHere = 1
			minEndingHere = 0
		} else {
			temp = maxEndingHere
			maxEndingHere = Max(minEndingHere * x[i], 1)
			minEndingHere = temp * x[i]
		}
		maxSoFar = Max(maxSoFar, maxEndingHere)
	}
	return maxSoFar
}

func main() {
	var (
		x = []int{2, -3, 0, -4, 3, -2, -3, 3}
		n = len(x) / 4
	)

	fmt.Println("Maximum Consecutive Elements Producr Program")
	fmt.Println("============================================")
	fmt.Print("Given Array :")
	fmt.Printf("%v\n", x)
	fmt.Printf("Maximum Product is %d\n", MaxProd(x, n))
}