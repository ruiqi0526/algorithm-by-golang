package main

import (
	"fmt"
)

const (
	YES = 1
	NO = 0
)

func Containment(left, right, mark []int, n int) {
	Sort(left, right, n)

	currentRightEnd := right[0] 
	mark[0] = NO
	for i := 1; i < n; i++ {
		if right[i] <= currentRightEnd {
			mark[i] = YES
		} else {
			mark[i] = NO
			currentRightEnd = right[i]
		}
	}
}

func Sort(left, right []int, n int) {
	Qsort(left, right, 0, n - 1)
}

func Qsort(left, right []int, first, last int) {
	var splitPoint int

	if first < last {
		Split(left, right, first, last, &splitPoint)
		if splitPoint - first < last - splitPoint {
			Qsort(left, right, first, splitPoint - 1)
			Qsort(left, right, splitPoint + 1, last)
		} else {
			Qsort(left, right, splitPoint + 1, last)
			Qsort(left, right, first, splitPoint - 1)
		}
	}
}

func Split(left, right []int, first, last int, sp *int) {
	x := left[first]
	y := right[first]
	currentSplit := first

	for i := first + 1; i <= last; i++ {
		if left[i] < x || (left[i] == x && right[i] > y) {
			currentSplit++
			Swap(left[currentSplit], left[i])
			Swap(right[currentSplit], right[i])
		}
	}
	Swap(left[first], left[currentSplit])
	Swap(right[first], right[currentSplit])
	*sp = currentSplit
}

func Swap(p, q int) {
	p, q = q, p
}

func main() {
	var (
		left = []int{0, 4, 7, 6, 4, 1, 3, 5, 0, 4}
		right = []int{3, 5, 10, 10, 7, 2, 5, 9, 4, 6}
		n = len(left) / 8
		mark = make([]int, 10)
		marker = []string{" ", "*"}
	)

	fmt.Println("Interval Containment Checking Program")
	fmt.Println("=====================================")

	Containment(left, right, mark, n)
	
	fmt.Println("Interval #   [left,right]")
	fmt.Println("----------   ------------")
	for i := 0; i < n; i++ {
		fmt.Printf("%d  [%d, %d]%s\n", i + 1, left[i], right[i], marker[mark[i]])
	}

	fmt.Printf("%s means the interval is contained in some other intervals\n", marker[1])
}