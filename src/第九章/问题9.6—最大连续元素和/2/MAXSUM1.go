package main

import (
	"fmt"
)

func Max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func MaxSum(x []int, left, right int) int {
	var (
		middle int
		maxToLeft int
		maxToRight int
		maxCrossing int
		sum int
		result int
	)

	if left > right {
		result = 0
	} else if left == right {
		result = func() int {
			if x[left] > 0 {
				return x[left]
			} else {
				return 0
			}
 		}()
	} else {
		middle = (left + right) / 2
		sum = 0
		maxToLeft = 0
		for i := middle; i >= left; i-- {
			sum += x[i]
			if maxToLeft < sum {
				maxToLeft = sum
			}
		}
		sum = 0
		maxToRight = 0
		for i := middle + 1; i <= right; i++ {
			sum += x[i]
			if maxToRight < sum {
				maxToRight = sum
			}
		}
		maxCrossing = maxToLeft + maxToRight
		if maxCrossing >= maxToLeft {
			result = Max(maxCrossing, maxToRight)
		} else {
			result = Max(maxToLeft, maxToRight)
		}
	}
	return result
}

func main() {
	var (
		x = []int{2, -3, 1, -1, 3, -2, -3, 3}
		n = len(x) / 4
	)

	fmt.Println("Maximum Consecutive Elements Sum Program")
	fmt.Println("========================================")
	fmt.Print("Given Array :")
	fmt.Printf("%v\n", x)
	fmt.Printf("Maximum Sum is %d\n", MaxSum(x, 0, n - 1))
}