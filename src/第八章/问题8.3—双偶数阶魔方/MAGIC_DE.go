package main

import (
	"fmt"
)

const (
	MAXSIZE = 20
	MARK = -1
)

func main() {
	var (
		square = make([][]int, MAXSIZE)
		n int
		count int
		invCount int
		marker int
		i, j int
	)
	for i, _ := range square {
		square[i] = make([]int, MAXSIZE)
	}

	fmt.Println("Doubly-Even Magic Square")
	fmt.Println("========================")
	fmt.Print("Order (4*m, m>0) please --> ")
	fmt.Scanln(&n)

	if n % 4 != 0 {
		fmt.Println("*** Illegal Order *****")
	} else {
		marker = MARK
		for i = 0; i < n / 2; i++ {
			for j = 0; j < n / 2; j++ {
				square[i][n - 1 - j] = marker
				square[i][j] = square[i][n - 1 - j]
				marker = -marker
			}
			marker = -marker
		}
		count = 1
		invCount = n * n
		for i = 0; i < n / 2; i++ {
			for j = 0; j < n; j++ {
				if square[i][j] != MARK {
					square[i][j] = count
					count++
					square[n - 1 - i][n - 1 - j] = invCount
					invCount--
				} else {
					square[i][j] = invCount
					invCount--
					square[n - 1 - i][n - 1 - j] = count
					count++
				}
			}
		}
		fmt.Println()
		for i = 0; i < n; i++ {
			for j = 0; j < n; j++ {
				fmt.Printf("%4d", square[i][j])
			}
			fmt.Println()
		}
	}
}