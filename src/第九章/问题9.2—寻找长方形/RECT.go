package main

import (
	"fmt"
)

func FindRectangle(len []int, n int) int {
	var (
		total int
		half int
		sum int
		start int
		end int
		count int
	)

	for i := 0; i < n; i++ {
		total += len[i]
	}
	half = total / 2
	sum = len[start]
	for start < n - 1 {
		if sum < half {
			start++
			sum += len[start]
		} else if sum > half {
			sum -= len[end]
			end++
		} else {
			start++
			sum += len[start]
			count++
		}
	}

	return count
}

func main() {
	var (
		x = []int{1, 3, 1, 1, 4, 1, 2, 1, 2, 2, 2}
		n = len(x) / 4
		y = []int{2, 2, 2, 3, 3}
		m = len(y) / 4
		z = []int{10, 2, 2, 2, 2}
		mn = len(z) / 4
	)

	fmt.Println("Rectangle Finding Program")
	fmt.Println("=========================")
	fmt.Print("First Arc Length Set  :")
	fmt.Printf("%v\n", x)
	fmt.Printf("There are %d pairs diametrical points.\n", FindRectangle(x, n))
	fmt.Print("Second Arc Length Set :")
	fmt.Printf("%v\n", y)
	fmt.Printf("There are %d pairs diametrical points.\n", FindRectangle(y, m))
	fmt.Print("Third  Arc Length Set :")
	fmt.Printf("%v\n", z)
	fmt.Printf("There are %d pairs diametrical points.\n", FindRectangle(z, mn))
}