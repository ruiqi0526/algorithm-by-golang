package main

import (
	"fmt"
)

func Min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func main() {
	var (
		money = make([]int, 101)
		base = []int{1, 3, 4}
		k = len(base) / 8
		n int
		i, j, MIN int
	)

	fmt.Println("Minimum Money Change Program")
	fmt.Println("============================")
	fmt.Print("Base Values : ")
	fmt.Println(base)
	fmt.Print("Your input please --> ")
	fmt.Scanln(&n)

	money[1] = 1
	for i = 2; i <= n; i++ {
		MIN = n
		for j = 0; j < k; j++ {
			if i >= base[j] {
				MIN = Min(money[i - base[j]] + 1, MIN)
			}
		}
		money[i] = MIN
	}
	fmt.Printf("Minimum = %d\n", money[n])
}