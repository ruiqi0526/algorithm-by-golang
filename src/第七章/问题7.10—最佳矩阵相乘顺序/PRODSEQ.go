package main

import (
	"fmt"
	"math"
)

func ProductSequence(dim []int, n int) {

	var (
		MinCost int
		Previous, ThisOne int
		pos, i, j, k int
	)

	cost := make([][]int, 8 * n)
	for i, _ := range cost {
		cost[i] = make([]int, 4 * n)
	}
	factor := make([][]int, 8 * n * n)
	for i, _ := range factor {
		factor[i] = make([]int, 4 * n * n)
	}

	for i = 0; i < n; i++ {
		cost[i][i] = 0
	}

	for j = 1; j < n; j++ {
		for i = j - 1; i >= 0; i-- {
			MinCost = math.MaxInt64
			pos = -1
			for k = i; k < j; k++ {
				Previous = cost[i][k] + cost[k + 1][j]
				ThisOne = dim[i] * dim[k + 1] * dim[j + 1]
				if MinCost > Previous + ThisOne {
					MinCost = Previous + ThisOne
					pos = k
				}
			}
			cost[i][j] = MinCost
			factor[i][j] = pos
		}
	}
	Display(factor, n, cost[0][n - 1])
}

func Display(factor [][]int, n, total int) {
	fmt.Print("The Factorization is ")
	ShowOrder(factor, 0, n-1)
	fmt.Printf("Number of Multiplications to be Perfomed = %d\n", total)
}

func ShowOrder(factor [][]int, start, end int) {
	if start == end {
		fmt.Printf("M%d", start)
	} else {
		fmt.Print("(")
		mid := factor[start][end]
		ShowOrder(factor, start, mid)
		fmt.Print("*")
		ShowOrder(factor, mid + 1, end)
		fmt.Print(")")
	}
}

func main() {
	var dim = []int{30, 1, 40, 10, 25, 1, 10, 30}
	var n = len(dim) / 4

	fmt.Println("Optimum Matrix Product Sequence")
    fmt.Println("===============================")
	fmt.Println("Given Dimension List :")
	for i := 0; i < n - 1; i++ {
          fmt.Printf("Matrix %2d has dimension %dx%d \n", i, dim[i], dim[i+1])
	}
    ProductSequence(dim, n - 1)
}