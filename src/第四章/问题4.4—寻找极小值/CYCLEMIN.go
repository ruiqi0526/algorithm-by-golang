package main

import "fmt"

func cyclic_min(x []int, n int) int {
	var (
		left int
		right int = n - 1
		mid int
	)

	for left < right {
		mid = (left + right) / 2
		if x[mid] < x[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	var x = []int{20, 23, 28, 35, 39, 40, 42, 8, 10, 15, 17, 19}
	var n int = len(x) / 8

	fmt.Println("Find Cyclic Minimum")
	fmt.Println("===================")
	fmt.Println("Given Array Sorted in Cyclic Fashion :")
	for i := 0; i < n; i++ {
		fmt.Printf("%d", x[i])
	}

	loc := cyclic_min(x, n)
	fmt.Printf("\nMinimum is located at x[%d] = %d\n", loc, x[loc])
}