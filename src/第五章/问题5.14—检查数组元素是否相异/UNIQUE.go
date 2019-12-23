package main

import (
	"fmt"
)

func unique(x []int, n int, number *int) {
	*number = 1
	for i := 1; i < n; i++ {
		if x[i] != x[i-1] {
			(*number)++
			x[*number] = x[i]
		}
	}
}

func main() {
	var x = []int{100, 37, 100, 37, 15, 111, 37, 15, 111, 98, 100, 98}
	var n = len(x) / 8
	var number int

	fmt.Println("Element Uniquness Program")
	fmt.Println("=========================")
	fmt.Println("Original Array  :")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", x[i])
	}

	unique(x, n, &number)
	fmt.Println()
	fmt.Println("Processed Array :")
	for i := 0; i < number; i++ {
		fmt.Printf("%d ", x[i])
	}
	fmt.Println()
}