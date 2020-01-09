package main

import (
	"fmt"
)

func readIn(connected [][]int, n int) {
	var i, j int

	fmt.Scanln(&i, &j)
	for i != 0 && j != 0 {
		connected[i - 1][j - 1] = 1
		connected[j - 1][i - 1] = 1
		fmt.Scanln(&i, &j)
	}
	for i = 0; i < n; i++ {
		connected[i][i] = 0
	}
}

func Hamilton(connected [][]int, cycle []int, n, start int) bool {
	var (
		top, i int
		visited = make([]int, 4 * n)
	)

	for i = 0; i < n; i++ {
		visited[i] = 0
	}
	visited[start] = 1

	cycle[0] = start
	cycle[1] = 0
	top = 1

	for {
		for i = cycle[top]; i < n; i++ {
			if connected[cycle[top - 1]][i] == 1 && visited[i] == 0 {
				break
			}
		}
		if i < n {
			cycle[top] = i
			visited[cycle[top]] = 1
			if top == n - 1 && connected[cycle[top]][start] == 1 {
				return true
			} else {
				top++
				cycle[top] = 0
			}
		} else {
			top--
			visited[cycle[top]] = 0
			if top == 0 {
				return false
			}
			cycle[top]++
		}
	}
}

func Display(cycle []int, n int) {
	fmt.Println("A Hamilton Cycle is Listed as Follows :")
	fmt.Printf("%d", cycle[0]+1)
	for i := 1; i < n; i++ {
		fmt.Printf("->%d", cycle[i] + 1)
	}
	fmt.Printf("->%d", cycle[0] + 1)
}

func main() {
	var (
		connected = make([][]int, 30)
		cycle = make([]int, 30)
		n int
	)
	for i, _ := range connected {
		connected[i] = make([]int, 30)
	}

	fmt.Println("Hamilton Cycle Program")
	fmt.Println("======================")
	fmt.Scanln(&n)
	readIn(connected, n)
	if Hamilton(connected, cycle, n, 0) == true {
		Display(cycle, n)
	} else {
		fmt.Println("NO Hamilton Cycle at all.")
	}
}