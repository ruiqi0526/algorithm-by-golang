package main

import "fmt"

type MATRIX [5][5]int

func matrix_search(mat MATRIX, n, key, index_x, index_y int) {
	i := 0
	j := n - 1

	index_x = -1
	index_y = -1
	for i < n && j >= 0 {
		if mat[i][j] < key {
			i++
		} else if mat[i][j] > key {
			j--
		} else {
			index_x = i
			index_y = j
		}
	}
}

func main() {
	var x  = MATRIX{{1, 3, 7, 9, 13},
					{2, 4, 8, 10, 14},
					{5, 6, 9, 12, 15},
					{7, 12, 13, 15, 19},
					{11, 13, 16, 18, 25}}

	var n int = 5
	var key, i, j int 

	fmt.Println("Special Format Matrix Search Program")
    fmt.Println("====================================")
	fmt.Println("Given Matrix :")
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			fmt.Printf("%d ", x[i][j])
		}
		fmt.Println()
	}

	fmt.Print("The Key to be Searched for --> ")
	fmt.Scanln(&key)
	matrix_search(x, n, key, i, j)
	if i < 0 || j < 0 {
		fmt.Println("NO! Key value not found")
	} else {
		fmt.Printf("The Key is located at mat[%d][%d]\n", i, j)
	}
}