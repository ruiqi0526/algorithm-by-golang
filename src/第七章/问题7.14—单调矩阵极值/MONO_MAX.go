package main

import (
	"fmt"
	"math"
)

func MonotoneMax(data [][]int, m, n int, loc []int) {
	GetMax(data, 0, m - 1, 0, n - 1, loc)
}

func GetMax(data [][]int, T, B, L, R int, loc []int) {
	var (
		midRow = (T + B) / 2
		pos, max int
		j int
	)

	if T <= B {
		max = math.MaxInt64
		for j = L; j <= R; j++ {
			if data[midRow][j] > max {
				max = data[midRow][j]
				pos = j
			}
		}
		loc[midRow] = pos
		GetMax(data, T, midRow - 1, L, pos, loc)
		GetMax(data, midRow + 1, B, pos, R, loc)
	}
}

func main() {
	var x = [][]int{
		{ 2, 5, 4, 3, 1, 2, 4},
		{ 1, 3, 2, 1, 0, 1, 2},
		{ 2, 4, 7, 3, 1, 3, 4},
		{ 5, 3, 1, 6, 4, 3, 2},
		{ 1, 3, 5, 7, 6, 5, 4},
		{ 2, 4, 6, 8, 6, 7, 9},
	}
	var location = make([]int, 10)
	var m = 6
	var n = 7

	fmt.Println("Monotone Matrix Max. Computation")
    fmt.Println("================================")
	fmt.Print("Input Matrix :")
	fmt.Printf("%v\n", x)

	MonotoneMax(x, m, n, location)
	fmt.Println("Row   Max.Col   Value")
	fmt.Println("---   -------   -----")
	for i := 0; i < m; i++ {
		fmt.Printf("%d%9d%9d\n", i, location[i], x[i][location[i]])
	}
}