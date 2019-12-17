package main

import (
	"fmt"
	"math/rand"
)

func median(x []int, n int) int {
	var (
		left = 0
		right = n - 1
		mid = (left + right) / 2
		split_point int
	)

	for {
		split(x, left, right, split_point)
		if split_point == mid {
			break
		} else if split_point > mid {
			right = split_point - 1
		} else {
			left = split_point + 1
		}
	}

	if n & 1 != 0 {
		return x[mid]
	} else {
		return (x[mid] + x[mid + 1]) / 2
	}
}

func SWAP(x, y int) {
	var t int
	t = x
	x = y
	y = t
}

func split(x []int, left, right, split_point int) {
	var split_data = x[left]
	split_point = left

	for i := left; i <= right; i++ {
		if x[i] < split_data {
			split_point++
			SWAP(x[split_point], x[i])
		}
	}
	SWAP(x[left], x[split_point])
}

func main() {
	var (
		x = make([]int, 20)
		n int
	)

	fmt.Println("Median Finder Program")
	fmt.Println("=====================")
	fmt.Print("How many data do you want to generate --> ")
	fmt.Scanln(&n)

	fmt.Println("Data Array : ")
	for i := 0; i < n; i++ {
		if i % 10 == 0 { fmt.Println() }
		x[i] = rand.Intn(100)
		fmt.Printf("%d ", x[i])
	}
	fmt.Printf("\nThe median is %d\n", median(x, n))
}