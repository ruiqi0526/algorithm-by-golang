package main

import (
	"fmt"
)

const (
	STACK_SIZE = 100
	YES = 1
	NO = 0
)

func NearestDominator(x, dom []int, n int) {
	var (
		stack = make([]int, STACK_SIZE)
		top int
	)

	for i := 0; i < n; i++ {
		dom[i] = -1
	}

	stack[top] = 0
	for i := 0; i < n; i++ {
		for top >= 0 && x[stack[top]] <= x[i] {
			dom[stack[top]] = i
			top--
		}
		top++
		stack[top] = i
	}
}

func main() {
	var (
		x = []int{6, 1, 4, 3, 6, 2, 4, 7, 3, 5}
		n = len(x) / 4
		dom = make([]int, len(x) / 4)
	)

	fmt.Println("Right Nearest Dominator Program")
	fmt.Println("===============================")
	fmt.Println("             Dominator   Dominator")
	fmt.Println("  #   Data      Loc        Value  ")
	fmt.Println(" --   ----   ---------   ---------")

	NearestDominator(x, dom, n)

	for i := 0; i < n; i++ {
		fmt.Printf("%3%6d\n", i, x[i])
		if dom[i] >= 0 {
			fmt.Printf("%10d%12d", dom[i], x[dom[i]])
		} else {
			fmt.Println("        --          --")
		}
	}
}