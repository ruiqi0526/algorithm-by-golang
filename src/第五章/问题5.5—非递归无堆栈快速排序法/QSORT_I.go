package qsort_i

import (
	"fmt"
)

const (
	YES int = 1
	NO int = 0
	ALWAYS int = 1
)

func SWAP(a, b int) {
	var t int

	t = a
	a = b
	b = t
}

func Sort(x []int, n int) {
	x[n] = -1
	for sorted := 0; sorted < n; sorted++ {
		for x[sorted] > 0 {
			key := x[sorted]
			split := sorted
			for next := sorted + 1; x[next] > 0; next++ {
				if x[next] < key {
					split++
					SWAP(x[split], x[next])
				}
			}
			SWAP(x[sorted], x[split])
			x[split] = -x[split]
		}
		x[sorted] = -x[sorted]
	}
}