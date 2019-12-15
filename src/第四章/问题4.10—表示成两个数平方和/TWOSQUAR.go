package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		given int
		row, column int
		count int
	)

	fmt.Println("Representing a Given Number as the Sum of Two Squares")
	fmt.Println("=====================================================")
	fmt.Print("An Integer Please ---> ")
	fmt.Scanln(&given)
	fmt.Println("Count      X      Y")
	fmt.Println("-----  -----  -----")

	row = 1
	column = int(math.Sqrt(float64(given)) + 0.5)
	count = 0
	for row <= given && column > 0 {
		if row * row + column * column == given {
			count++
			fmt.Printf("%d%d%d\n", count, row, column)
			column++
		} else if row * row + column * column > given {
			column++
		} else {
			row++
		}
	}

	if count == 0 {
		fmt.Println("Sorry, NO ANSWER found.")
	} else {
		fmt.Println("There are %d possible answers.", count)
	}

}