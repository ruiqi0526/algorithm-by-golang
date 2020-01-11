package main

import (
	"fmt"
)

func MaxSum(x []int, n int) int {
	var (
		maxEndingHere int
		maxSoFar int
	)

	for i := 0; i < n; i++ {
		if maxEndingHere + x[i] < 0 {
			maxEndingHere = 0
		} else {
			maxEndingHere += x[i]
		}
		if maxEndingHere > maxSoFar {
			maxSoFar = maxEndingHere
		}
	}
	return maxSoFar
}

func main() {
	var (
		x = []int{2, -3, 1, -1, 3, -2, -3, 3}
		n = len(x) / 4
	)

	fmt.Println("Maximum Consecutive Elements Sum Program");
	fmt.Println("========================================");
	fmt.Print("Given Array :");
	fmt.Printf("%v\n", x)
	fmt.Printf("Maximum Sum is %d\n", MaxSum(x, n))
}