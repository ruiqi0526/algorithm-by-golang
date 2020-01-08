package main

import (
	"fmt"
)

func main() {
	var (
		count int
		row int
		column int
		order int
		matrix = make([][]int, 20)
	)

	for i, _ := range matrix {
		matrix[i] = make([]int, 20)
	}

	fmt.Println("Odd Order Magic Square Generator")
	fmt.Println("================================")
	fmt.Print("Order Please --> ")
	fmt.Scanln(&order)

	if order > 20 {
		fmt.Printf("*** ERROR ***  Order should be <= %d\n", 20)
	} else if order % 2 == 0 {
		fmt.Println("*** ERROR ***  Order must be an odd integer")
	} else {
		column = order / 2
		for count = 1; count <= order * order; count++ {
			matrix[row][column] = count
			if count % order == 0 {
				row++
			} else {
				row = func() int {
					if row == 0 {
						return order - 1
					} else {
						return row - 1
					}
				}()
				column = func() int {
					if column == order - 1 {
						return 0
					} else {
						return column + 1
					}
				}()
			}
		}
		fmt.Printf("Magic Square of order %d :\n", order)
		for row = 0; row < order; row++ {
			for column = 0; column < order; column++ {
				fmt.Printf("%4d", matrix[row][column])
			}
			fmt.Println()
		}
	}
}