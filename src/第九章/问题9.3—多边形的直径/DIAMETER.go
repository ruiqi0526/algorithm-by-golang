package main

import (
	"fmt"
	"math"
)

func Diameter(len []int, n int, first, last *int) {
	var (
		start, end int
		lead, follow int
		diff, absDiff int
		currentMin int
	)

	for i := 0; i < n; i++ {
		diff -= len[i]
	}

	absDiff = -diff
	currentMin = -diff
	for lead < n {
		if currentMin >= absDiff {
			currentMin = absDiff
			start = lead
			end = follow
		}
		if diff >= 0 {
			diff -= len[follow] * 2
			follow++
		} else {
			diff += len[lead] * 2
			lead++
		}
		absDiff = int(math.Abs(float64(diff)))
	}
	*first = end
	*last = start
}

func main() {
	var (
		x = []int{1, 2, 2, 1, 3, 2, 4, 1, 1, 2, 2, 3}
		n = len(x) / 4
		start, end int
	)

	fmt.Println("Polygon Diameter Computation")
	fmt.Println("============================")
	fmt.Println("Side #     Length")
	fmt.Println("-----------------")
	for i := 0; i < n; i++ {
		fmt.Printf("%5d%11d\n", i, x[i])
	} 
	Diameter(x, n, &start, &end)
	fmt.Printf("Diameter is generated from point %d to %d\n", start, end)
}