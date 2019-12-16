package qsort

import (
	"fmt"
)

func Sort(input []int, n int) {
	var (
		first = 0
		last = n - 1
	)

	qsort(input, first, last)
}

func qsort(input []int, first, last int) {
	var split_point int

	if first < last {
		split(input, first, last, split_point)
		if split_point - first < last - split_point {
			qsort(input, first, split_point - 1)
            qsort(input, split_point+1, last)
		} else {
			qsort(input, split_point+1, last)
            qsort(input, first, split_point - 1)
		}
	}
}

func split(input []int, first, last int, split_point int) {
	x := input[first]
	current_split := first
	for unknown := first + 1; unknown <= last; unknown++ {
		if input[unknown] < x {
			current_split++
			swap(input[current_split], input[unknown])
		}
	}
	swap(input[first], input[current_split])
	split_point = current_split
}

func swap(a, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}