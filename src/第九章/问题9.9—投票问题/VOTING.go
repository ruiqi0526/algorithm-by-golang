package main

import (
	"fmt"
)

func Majority(vote []int, n int) int {
	var (
		candidate int
		voteCount int
		count int
		answer int
	)

	candidate = vote[0]
	voteCount = 1
	for i := 0; i < n; i++ {
		if voteCount == 0 {
			candidate = vote[i]
			voteCount = 1
		} else if candidate == vote[i] {
			voteCount++
		} else {
			voteCount--
		}
	}
	if voteCount == 0 {
		answer = 0
	} else {
		for i := 0; i < n; i++ {
			if vote[i] == candidate {
				count++
			}
		}
		answer = func() int {
			if int(float64(n) / 2.0 + 0.5) == 1 {
				return candidate
			} else {
				return 0
			}
		}()
	}
	return answer
}

func main() {
	var (
		x = []int{2, 2, 4, 2, 1, 2, 5, 2, 2, 8}
		n = len(x) / 4
	)

	fmt.Println("Majority Counting Program")
	fmt.Println("=========================")
	fmt.Println("  No   Vote")
	fmt.Println("  --   ----")
	for i := 0; i < n; i++ {
		fmt.Printf("%4d%6d\n", i, x[i])
	}
	answer := Majority(x, n)
	if answer > 0 {
		fmt.Printf("Majority is %d\n", answer)
	} else {
		fmt.Println("There is no majority\n")
	}
}