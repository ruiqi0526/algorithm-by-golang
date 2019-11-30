package main

import (
	"fmt"
)

func dominance_count(f []int, g []int, m int, n int) int {
	var (
		index_f int = 0
		index_g int
		count int
	)

	for index_f < m && index_g < n {
		if f[index_f] <= g[index_g] {
			index_f++
		} else {
			index_g++
			count += m - index_f
		}
	}

	return count
}


func main() {
	x := []int{1, 3, 5, 7, 9}
	nx := len(x)

	y := []int{2, 3, 4, 7, 8}
	ny := len(y)

	fmt.Println("Array1 is :", x)
	fmt.Println("Array2 is :", y)
	fmt.Printf("There are %d dominance pairs.\n", dominance_count(x, y, nx, ny))
}