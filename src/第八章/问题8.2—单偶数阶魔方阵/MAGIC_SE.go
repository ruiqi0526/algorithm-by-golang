package main

import (
	"fmt"
)

const MAXSIZE = 30

func SinglyEven(matrix [][]int, n int) {
	half := n / 2

	MagicO(matrix, half)
	Exchange(matrix, n)
}

func MagicO(matrix [][]int, n int) {
	row := 0
	column := n / 2
	for count := 1; count <= n * n; count++ {
		matrix[row][column] = count
		matrix[row + n][column + n] = count + n * n
		matrix[row][column + n] = count + n * n * 2
		matrix[row + n][column] = count + n * n * 3
		if count % n == 0 {
			row++
		} else {
			row = func() int {
				if row == 0 {
					return n - 1
				} else {
					return row - 1
				}
			}()
			column = func() int {
				if column == n - 1 {
					return 0
				} else {
					return column + 1
				}
			}()
		}
	}
}

func SWAP(a, b int) {
	a, b = b, a
}

func Exchange(x [][]int, n int) {
	var (
		width = n / 4
		width1 = width - 1
		i, j int
	)

	for i = 0; i < n / 2; i++ {
		if i != width {
			for j = 0; j < width; j++ {
				SWAP(x[i][j], x[n / 2 + i][j])
			}
			for j = 0; j < width1; j++ {
				SWAP(x[i][n - 1 - j], x[n / 2 + i][n - 1 - j])
			}
		} else {
			for j = 0; j <= width1; j++ {
				SWAP(x[width][j], x[n / 2 + width][j])
			}
			for j = 0; j < width1; j++ {
				SWAP(x[width][n - 1 - j], x[n / 2 + width][n - 1 - j])
			}
		}
	}
}

func main() {
	var matrix = make([][]int, MAXSIZE)
	for i, _ := range matrix {
		matrix[i] = make([]int, MAXSIZE)
	}

	var n int

	fmt.Println("ingly-Even Order Magic Square")
	fmt.Println("=============================")
	fmt.Print("rder Please (must be 2*(2k+1)) --> ")
	fmt.Scanln(&n)

	if n % 2 == 0 && (n / 2) % 2 == 1 {
		fmt.Println()
		SinglyEven(matrix, n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Printf("%4d", matrix[i][j])
			}
			fmt.Println()
		}
	} else {
		fmt.Println("*** Illegal Order ***")
	}
}