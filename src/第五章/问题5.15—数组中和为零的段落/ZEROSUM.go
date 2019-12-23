package main

import (
	"fmt"
)

func zero_sum(x []int, n int) int {
	var i int

	for i = 1; i < n; i++ {
		x[i] += x[i - 1]
	}
	for i = 1; i < n && x[i] != x[i - 1]; i++ {}

	if i == n {
		return 0
	} else {
		return 1
	}
	
}

func main() {
	var x = []int{4, -1, 2, 1, -2, -1, 5}
	var n = len(x) / 8

	fmt.Println("Zero Sum Sunarray Check Program")
	fmt.Println("===============================")
	fmt.Println("Given Array :")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", x[i])
	}
	if zero_sum(x, n) == 1 {
		fmt.Println("YES, there is some range summing up to 0")
	} else {
		fmt.Println("NO! no such subarray found")
	}
}