package main

import (
	"fmt"
)

func longest_plateau(x []int, n int) int {
	length := 1

	for i := 1; i < n; i++ {
		if x[i] == x[i - length] {
			length++
		}
	}

	return length
}


func main() {

	var x = []int{1,2,2,3,3,3,4,5,5,6}
	n := len(x)

	fmt.Print("The given array:")
	fmt.Println(x)
	fmt.Printf("Length of the longest plateau is %d\n", longest_plateau(x, n))
	
}