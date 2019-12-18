package main

import (
	"fmt"
	"math/rand"
)

func dup_sort(x []int, n int) {
	var (
		dup int
		count int
		i, j int
	)

	for i = 1; i < n; i++ {
		if x[i] == x[i - 1] {
			dup = x[i]
			break
		}
	}
	count = i
	for i = 0; i < n; i++ {
		if x[i] != dup {
			x[count] = x[i]
			count++
		}
	}
	x[count] = dup
	count++

	index := n - 1
	for i = count - 1; i >= 0 && x[i] != dup; i-- {
		x[index] = x[i]
		index--
	}

	if i >= 0 && x[i] == dup {
		for j = i + 1; j <= index; j++ {
			x[j] = dup
		}
	}
}

func main() {
	var x = make([]int, 100)

	dup := rand.Intn(100)
	for i := 0; i < 100; i++ {
		x[i] = dup
	}

	for i := 1; i <= 10; i++ {
		j := rand.Intn(100)
		k := rand.Intn(100)
		x[j] = k
	}

	fmt.Println("Soring with a lot of Duplicated Data")
	fmt.Println("====================================")
	fmt.Println("Generated Data :")
	for i := 0; i < 100; i++ {
		if i % 10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%d ", x[i])
	}

	dup_sort(x, 100)
	fmt.Println("Sorted Data :")
	for i := 0; i < 100; i++ {
		if i % 10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%d ", x[i])
	}
}