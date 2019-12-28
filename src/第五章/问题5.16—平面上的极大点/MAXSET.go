package main

import (
	"fmt"
	"math"
)

func max_set(x, xx, y, yy []int, n int, no *int) {
	max := math.MinInt64
	*no = 0

	for i := n - 1; i >= 0; i-- {
		if y[i] > max {
			xx[*no] = y[i]
			yy[*no] = y[i]
			(*no)++
			max = y[i]
		}
	}
}

func sort(x, y []int, n int) {
	qsort(x, y, 0, n - 1)
}

func qsort(x, y []int, first, last int) {
	var split_pt int

	if first < last {
		split(x, y, first, last, &split_pt)
		if split_pt - first < last - split_pt {
			qsort(x, y, first, split_pt - 1)
			qsort(x, y, split_pt + 1, last)
		} else {
			qsort(x, y, split_pt + 1, last)
			qsort(x, y, first, split_pt - 1)
		}
	}
}

func split(x, y []int, first, last int, split_pt *int) {
	xx := x[first]
	current_split := first

	for unknown := first + 1; unknown <= last; unknown++ {
		if x[unknown] < xx {
			current_split++
			swap(x[current_split], x[unknown])
			swap(y[current_split], y[unknown])
		}
	}
	swap(x[first], x[current_split])
	swap(y[first], y[current_split])
	*split_pt = current_split
}

func swap(x, y int) {
	y, x = x, y
}

func main() {
	var (
		x = []int{7, 3, 1, 8, 5, 4, 10, 9, 6, 0}
		y = []int{5, 6, 10, 7, 8, 9, 6, 8, 9, 8}
		n = len(x) / 8
		dom_x = make([]int, 100)
		dom_y = make([]int, 100)
		count int
	)

	fmt.Println("Maximal Point Set Program")
	fmt.Println("=========================")
	fmt.Println("Given Points :")
	fmt.Print("x[] = ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", x[i])
	}
	fmt.Print("\ny[] = ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", y[i])
	}

	max_set(x, dom_x, y, dom_y, n, &count)
	fmt.Printf("\nThere are %d Maximal Points :", count);
	fmt.Print("\nx[] = ")
	for i := 0; i < count; i++ {
		fmt.Printf("%d ", dom_x[i])
	}
	fmt.Print("\ny[] = ")
	for i := 0; i < count; i++ {
		fmt.Printf("%d ", dom_y[i])
	}
}