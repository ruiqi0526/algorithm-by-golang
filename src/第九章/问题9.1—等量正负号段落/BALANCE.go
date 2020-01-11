package main

import (
	"fmt"
)

func Max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func PosNegBalance(x []int, n int) int {
	var (
		length int
		negOver int
		minOver int
		maxOver int
		mid int
		loc = make([]int, 4 * (n + n))
	)

	mid = n
	for i := 0; i < n; i++ {
		if x[i] < 0 {
			negOver++
		} else if x[i] > 0 {
			negOver--
		}
		if negOver < minOver {
			minOver = negOver
			loc[mid + minOver] = i + 1
		} else if negOver > maxOver {
			maxOver = negOver
			loc[mid + maxOver] = i + 1
		}
		length = Max(length, i + 1 - loc[mid + minOver])
	}

	return length
}

func main() {
	var (
		x = []int{1, 1, -1, -2, 0, 1, 3, -1, 2, -1}
		n = len(x) / 4
	)

	fmt.Println("Maximum Interval of Balanced Positive/Negative Numbers")
	fmt.Println("======================================================")
	fmt.Print("Given Array :")
	fmt.Printf("%v\n", x)
	fmt.Printf("Interval Length = %d\n", PosNegBalance(x, n))
}