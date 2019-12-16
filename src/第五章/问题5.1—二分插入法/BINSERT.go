package binsert

import "fmt"

var (
	YES int = 1
	NO int = 0
	SIZE int = 8
)

func Sort(input []int, n int) {

	for current := 1; current < n; current++ {
		x := input[current]
		pos := -1
		if x < input[0] {
			pos = 0
		} else if x <= input[current - 1] {
			low := 0
			high := current - 1
			for high - low > 1 {
				mid := (high + low) / 2
				if x >= input[mid] {
					low = mid
				} else {
					high = mid
				}
			}
			pos = low + 1
		}
		if pos >= 0 {
			input[pos + 1] = input[pos][0:SIZE * (current - pos) + 1]
			input[pos] = x
		}
	}
}