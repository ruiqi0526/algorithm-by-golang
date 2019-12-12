package main

import "fmt"

func index_search(x [8]int, n int) int {
	var (
		first int
		middle int
		last int = n - 1
	)

	index := -1
	for first <= last {
		middle = (first + last) / 2
		if x[middle] == middle {
			index = middle
			break
		} else if x[middle] > middle {
			last = middle - 1
		} else {
			first = middle + 1
		}
	} 
	return index
}

func main() {
	var x = [8]int{-1, 0, 1, 3, 5, 7, 9, 10}
	var n = len(x) / 8
	var answer int

	fmt.Println("Index Search Program")
	fmt.Println("====================")
	fmt.Print("Given Array :")
	for i := 0; i < n; i++ {
		fmt.Printf("%d", x[i])
	}

	answer = index_search(x, n)
	if answer >= 0 {
		fmt.Println("YES, x[%d] = %d has been found.", answer, answer)
	} else {
		fmt.Println("NO, there is no element with x[i] = i")
	}
}