package  main

import "fmt"

func coincidence_count(f []int, g []int, m int, n int) int {
	var (
		index_f int = 0
		index_g int
		count int
	)

	for index_f < m && index_g < n {
		if f[index_f] < g[index_g] {
			index_f++
		}else if f[index_f] > g[index_g] {
			index_g++
		}else{
			count++
			index_g++
			index_f++
		}
	}
	
	return count
}

func main() {
	x := []int{1, 3, 4, 7, 9}
	nx := len(x)

	y := []int{3, 5, 7, 8, 10}
	ny := len(y)

	fmt.Println("Array1 is :", x)
	fmt.Println("Array2 is :", y)
	fmt.Println("There are %d equal numbers.", coincidence_count(x, y, nx, ny))
}