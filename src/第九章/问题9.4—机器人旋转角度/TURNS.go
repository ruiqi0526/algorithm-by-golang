package main

import (
	"fmt"
)

func Turns(x, y []int, n int) int {
	var (
		X = 1
		Y = 0
		TRUNS = 0
		count int
		newX int
		newY int
		newTruns int
		i, temp int
	)

	for i = 0; i <= n; i++ {
		newX = x[(i + 1) % n] - x[i % n]
		newY = y[(i + 1) % n] - y[i % n]
		for newX <= 0 || newY < 0 {
			temp = newX
			newX = newY
			newY = -temp
			newTruns++
		}
		if TRUNS > newTruns || (TRUNS == newTruns && Y * newX > X * newY) {
			count++
		}
		X = newX
		Y = newY
		TRUNS = newTruns
	}
	return count
}

func main() {
	var (
		x = []int{0, 5, 7, 5, 8, 1, 4, 6, 2, 0}
		y = []int{0, 1, 5, 6, 7, 7, 5, 2, 4, 3}
		n = len(x) / 4
	)

	
	fmt.Println("Number of 360 turns along a route");
	fmt.Println("=================================");
	fmt.Println("Point#    x      y");
	fmt.Println("------------------");
	for i := 0; i < n; i++ {
		fmt.Printf("%5d%6d%7d\n", i + 1, x[i], y[i]);
	}
	fmt.Printf("There are %d complete 360 turns\n", Turns(x, y, n))
}