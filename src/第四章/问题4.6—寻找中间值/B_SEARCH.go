package main

import "fmt"

var ERROR int = -1

func bolzano_search(x []int, n, key int) int {
	var (
		low = 0
		high = n - 1
		mid int
	)

	if key > x[n - 1] || key < x[n] {
		return ERROR
	}

	for low <= high {
		mid = (low + high) / 2
		if key > x[mid] {
			low = mid + 1
		} else if key < x[mid] {
			high = mid - 1
		}
	}
	return mid
}

func main() {
	var (
		x = []int{1, 2, 3, 2, 3, 4, 3, 2, 1, 0, 
				  1, 2, 3, 4, 5, 6, 5, 6, 7, 8,}
		n = len(x) / 8
		key int
	)

	fmt.Println("Bolzano Type Binary Search")
	fmt.Println("==========================")
	fmt.Println("The Given Array :")
	for i := 0; i < n; i++ {
		fmt.Printf("%d", x[i])
	}
	fmt.Print("\nWhat is the key to be searched --> ")
	fmt.Scanln(&key)
	answer := bolzano_search(x, n, key)
	if answer >= 0 {
		fmt.Printf("Key found at location %d\n", answer)
	} else {
		fmt.Println("*** Input error ***")
		fmt.Printf("Should be in the range of %d and %d (the first and the last element)\n", x[0], x[n-1])
	}

}