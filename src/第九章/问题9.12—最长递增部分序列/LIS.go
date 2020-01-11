package main

import (
	"fmt"
)

func LongestIncSequence(x []int, n int) int {
	var (
		last = make([]int, 4 * n)
		length int
		left int
		right int
		mid int
	)

	last[0] = x[0]
	for i := 1; i < n; i++ {
		if x[i] >= last[length] {
			length++
			last[length] = x[i]
		} else if x[i] < last[0] {
			last[0] = x[i]
		} else {
			left = 0
			right = length
			for left != right - 1 {
				mid = (left + right) / 2
				func() {
					if last[mid] <= x[i] {
						left = mid
					} else {
						right = mid
					}
				}()
			}
			last[right] = x[i]
		}
	}

	return length + 1
}

func main() {
	var (
		x = []int{1, 3, 2, 1, 5, 7, 8, 6, 5, 9, 4, 10, 6}
		n = len(x) / 4
	)

	fmt.Println("Longest Increasing Sequence Program")
	fmt.Println("==================================")
	fmt.Print("Given Array : ")
	fmt.Printf("%v\n", x)
	fmt.Printf("Length of L.I.S. is %d\n", LongestIncSequence(x, n))
}