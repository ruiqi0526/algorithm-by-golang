package main

import "fmt"

func minOfTwo(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minOfThree(a, b, c int) int {
	if a < b {
		return minOfTwo(a, c)
	} else {
		return minOfTwo(b, c)
	}
}

func maxSquare(x [10][10]int, n, key, row, column, width int) {
	var size [10][10]int

	for j := 0; j < n; j++ {
		size[n - 1][j] = x[n - 1][j]
	}

	for i := n - 2; i >= 0; i-- {
		size[i][n - 1] = x[i][n - 1]
		for j := n - 2; j >= 0; j-- {
			size[i][j] = func() int {
				if x[i][j] == key {
					return 1 + minOfThree(size[i + 1][j], size[i][j + 1], size[i + 1][j + 1])
				} else {
					return 0
				}
			}()
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if size[i][j] > width {
				width = size[i][j]
				row = i
				column = j
			}
		}
	}
}

func main() {
	var x = [10][10]int{
		{ 1, 1, 1, 0, 0, 0, 0, 0},
		{ 0, 1, 1, 1, 1, 0, 0, 0},
		{ 0, 0, 1, 1, 1, 1, 1, 0},
		{ 0, 0, 0, 1, 1, 1, 1, 1},
		{ 0, 1, 1, 1, 1, 1, 1, 1},
		{ 0, 1, 1, 1, 1, 0, 1, 1},
		{ 0, 0, 1, 1, 1, 0, 0, 0},
		{ 0, 0, 1, 1, 1, 1, 0, 0},
	}

	var n int = 8
	var (
		row int = 1
		column int
		width int
		key int
	)

	fmt.Println("Find Maximum Size Square of Same Data");
    fmt.Println("=====================================");
	fmt.Println("Given Matrix : ")

	for row = 0; row < n; row++ {
		for column = 0; column < n; column++ {
			fmt.Printf("%d ", x[row][column])
		}
		fmt.Println()
	}
	fmt.Printf("Key to be searched = %d\n", key)
	maxSquare(x, n, key, row, column, width)
	fmt.Printf("Maximum Square is located at x[%d][%d] with size %d\n", row, column, width)
}