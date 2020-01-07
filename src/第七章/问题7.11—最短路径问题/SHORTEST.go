package main

import (
	"fmt"
	"math"
)

const (
	MAXSIZE = 20
)

func MaxSum(a, b int) int {
	if a != math.MaxInt64 && b != math.MaxInt64 {
		return a + b 
	} else {
		return math.MaxInt64
	}
}

func Floyd(dist [][]int, path [][]int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			path[i][j] = i
		}
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][j] > MaxSum(dist[i][k], dist[k][j]) {
					path[i][j] = path[k][j]
					dist[i][j] = MaxSum(dist[i][k], dist[k][j])
				}
			}
		}
	}
}

func DisplayPath(dist [][]int, path [][]int, n int) {
	var (
		i, j, k int
		count int
		chain = make([]int, 4 * n)
	)

	fmt.Println("Origin->Dest    Dist    Path")
	fmt.Println("------------    ----    ----")	

	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			if i != j {
				fmt.Printf("%d -> %d", i + 1, j + 1)
				if dist[i][j] == math.MaxInt64 {
					fmt.Print("  NA    ")
				} else {
					fmt.Printf("%d   ", dist[i][j])
					k = j
					for i != k {
						chain[count] = path[i][k]
						k = chain[count]
						count++
					}
					Reverse(chain, count)
					fmt.Printf("%d", chain[0] + 1)
					for k = 1; k < count; k++ {
						fmt.Printf("->%d", chain[k] + 1)
					}
					fmt.Printf("->%d", j + 1)
				}
			}
		}
	}
}

func Swap(i, j int) {
	i, j = j, i
}

func Reverse(x []int, n int) {
	i := 0

	for j := n - 1; i < j; j-- {
		Swap(x[i], x[j])
		i++
	}
}

func Readin(dist [][]int, number *int) {
	var (
		origin, dest, length, n int
		i, j int
	)

	fmt.Scanln(&n)
	*number = n
	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			dist[i][j] = math.MaxInt64
		}
		dist[i][i] = 0
	}

	fmt.Scanln(&origin, &dest, &length)
	for origin != 0 && dest != 0 && length != 0 {
		dist[origin + 1][dest + 1] = length
		fmt.Scanln(&origin, &dest, &length)
	}
}

func main() {
	var (
		dist = make([][]int, 20)
		path = make([][]int, 20)
		n int
	)

	for i, _ := range dist {
		dist[i] = make([]int, 20)
	}
	for i, _ := range path {
		path[i] = make([]int, 20)
	}

	fmt.Println("All-Shortest Program")
	fmt.Println("====================")
	
	Readin(dist, &n)
	Floyd(dist, path, n)
	DisplayPath(dist, path, n)
}