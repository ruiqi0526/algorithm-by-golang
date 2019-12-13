package main

import "fmt"
import "math"

func inf_search(x []int, GIVEN int) int {
	var left, right int

	bound(x, GIVEN, left, right)
	return search(x, GIVEN, left, right)
}

func bound(x []int, GIVEN int, start, end int) {
	var delta = 1

	start = 1
	end = start + delta
	for !(x[start] <= GIVEN && GIVEN < x[end]) {
		delta += delta
		start = end
		end = start + delta 
	}
	end--
}

func search(x []int, GIVEN int, low, high int) int {
	
	for low <= high {
		mid := (low + high) / 2
		if GIVEN < x[mid] {
			high = mid - 1
		} else if GIVEN > x[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	var number = []int{ 1,   3,   5,   8,  10, 
						21,  25,  50,  67,  68,
						70,  84, 100, 130, 150, 
	   					151, 170, 180, 200, 300,
	   					459, 480, 499, 503, 555,
						570, 623, 699, 784, 981,}
	
	var input = make([]int,100)
	var key int

	for i := 0; i < 30; i++ {
		input[i] = number[i]
	}
	for i := 30; i < 100; i++ {
		input[i] = math.MaxInt64
	}

	fmt.Println("Infinite Search Program")
    fmt.Println("=======================")
	fmt.Print("Given Infinite Sorted Array :")
	for i := 0; i < 100; i++ {
		if i % 15 == 0 {
			fmt.Println()
		}
		if input[i] < math.MaxInt64 {
			fmt.Printf("%d ", input[i])
		} else {
			fmt.Print(" inf")
		}
	}

	fmt.Print("\nYour Search Key Please --> ")
	fmt.Scanln(&key)

	answer := inf_search(input, key)
	if answer >= 0 {
		fmt.Printf("Key found at position %d\n", answer)
	} else {
		fmt.Println("Key NOT FOUND!")
	}
						   
}