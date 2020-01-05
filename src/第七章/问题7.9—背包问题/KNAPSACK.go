package main

import (
	"fmt"
)

const (
	YES = 1
	NO = 0
	SIZE = 27
	NOT_FOUND = -1
)

func Knapsack(size []int, n, SIZE int, result []int) int {
	var (
		count int
		i, j int
	)

	exist := make([][]int, 8 * (n + 1))
	for i, _ := range exist {
		exist[i] = make([]int, 8 * (n + 1) * (SIZE + 1))
	}
	member := make([][]int, 8 * (n + 1))
	for i, _ := range member {
		member[i] = make([]int, 8 * (n + 1) * (SIZE + 1))
	}

	exist[0][0] = YES
	for j = 1; j <= SIZE; j++ {
		exist[0][j] = NO
	}

	for i = 1; i <= n; i++ {
		for j = 0; j <= SIZE; j++ {
			member[i][j] = NO
			exist[i][j] = member[i][j] 
		}
		if exist[i - 1][j] == YES {
			exist[i][j] = YES
			member[i][j] = NO
		} else if j >= size[i - 1] {
			exist[i - 1][j - size[i - 1]] = YES
			member[i][j] = YES
		}
	}

	if exist[n][SIZE] == YES {
		j = SIZE
		for i = n; i != 0 && j != 0; {
			if member[i][j] == YES {
				i--
				result[count] = i
				count++
				j -= size[i]
			} else {
				i--
			}
		}
		Reverse(result, count)
	}

	return count
}

func SWAP(a, b int) {
	a, b = b, a
}

func Reverse(x []int, n int) {
	j := n - 1
	for i := 0; i < j; i++ {
		SWAP(x[i], x[j])
		j--
	}
}

func main() {
	var (
		size = []int{1, 2, 4, 7, 10, 12, 13, 15, 17}
		n = len(size) / 8
		count, i int
		result = make([]int, 5)
	)

	fmt.Println("Knapsack Problem")
	fmt.Println("================")
	fmt.Print("Given Item Size Array : ")
	fmt.Printf("%v\n", size)
	fmt.Printf("Given Knapsack Size = %d\n", SIZE)

	count = Knapsack(size, n, SIZE, result)
	if count == NOT_FOUND {
		fmt.Println("No Solution Found.")
	} else {
		fmt.Printf("Solution is %d = %d\n", SIZE, size[result[0]])
		for i = 1; i < count; i++ {
			fmt.Printf("%d\n", size[result[i]])
		}
	}
}