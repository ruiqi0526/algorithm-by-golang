package main

import (
	"fmt"
	"math"
)

func min(x int, y int) int {
	if x < y {
		return x
	}else{
		return y
	}
}

func min_distance(x []int, y []int, m int, n int) int {
	var (
		mininum int = math.MaxInt64
		index_x int = 0
		index_y int 
	)

	for index_x < m && index_y < n {
		if x[index_x] >= y[index_y] {
			mininum = min(mininum, x[index_x] - y[index_y])
			index_y++
		}else{
			mininum = min(mininum, y[index_y] - x[index_x])
			index_x++
		}
	}

	return mininum
}

func main() {
	x := []int{1, 3, 5, 7, 9}
	m := len(x)

	y := []int{2, 6, 8}
	n := len(y)

	fmt.Println("Given Array1:", x)
	fmt.Println("Given Array2:", y)
	fmt.Printf("Mininum Distance = %d\n",min_distance(x, y, m, n))
}