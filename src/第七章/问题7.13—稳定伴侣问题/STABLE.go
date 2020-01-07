package main

import (
	"fmt"
)

const (
	SIZE = 10
	FREE = -1
)

func StableMarriage(man [][]int, woman [][]int, n int, mEng, wEng []int) {
	var (
		boy, girl, top int
		rank = make([][]int, 8 * n)
		waiting = make([]int, 8 * n)
		next = make([]int, 8 * n)
	)

	for i, _ := range rank {
		rank[i] = make([]int, 8 * n * n)
	} 

	GetMem(woman, n, rank, waiting, next)
	for girl = 0; girl < n; girl++ {
		wEng[girl] = FREE
	}

	for top = n - 1; top >= 0; {
		boy = waiting[top]
		next[top]++
		girl = man[boy][next[boy]]
		if wEng[girl] == FREE {
			wEng[girl] = boy
			mEng[boy] = girl
			top--
		} else if rank[girl][boy] < rank[girl][wEng[girl]] {
			waiting[top] = wEng[girl]
			wEng[girl] = boy
			mEng[boy] = girl
		}
	}
}

func GetMem(w [][]int, n int, rank [][]int, wait, next []int) {
	for i := 0; i < n; i++ {
		wait[i] = n - 1 - i
		next[i] = 0
		for j := 0; j < n; j++ {
			rank[i][w[i][j]] = j
		}
	}
}

func main() {
	var man = [][]int{
		{ 4, 6, 0, 1, 5, 7, 3, 2},
		{ 1, 2, 6, 4, 3, 0, 7, 5},
		{ 7, 4, 0, 3, 5, 1, 2, 6},
		{ 2, 1, 6, 3, 0, 5, 7, 4},
		{ 6, 1, 4, 0, 2, 5, 7, 3},
		{ 0, 5, 6, 4, 7, 3, 1, 2},
		{ 1, 4, 6, 5, 2, 3, 7, 0},
		{ 2, 7, 3, 4, 6, 1, 5, 0},
	}
	var woman = [][]int{
		{ 4, 2, 6, 5, 0, 1, 7, 3},
		{ 7, 5, 2, 4, 6, 1, 0, 3},
		{ 0, 4, 5, 1, 3, 7, 6, 2},
		{ 7, 6, 2, 1, 3, 0, 4, 5},
		{ 5, 3, 6, 2, 7, 0, 1, 4},
		{ 1, 7, 4, 2, 3, 5, 6, 0},
		{ 6, 4, 1, 0, 7, 5, 3, 2},
		{ 6, 3, 0, 4, 1, 2, 5, 7},
	}
	var mEng = make([]int, SIZE)
	var nEng = make([]int, SIZE)
	var n = 8

	fmt.Println("Stable Marriage Problem")
	fmt.Println("=======================")
	fmt.Println("Man's Preference Matrix")
	fmt.Printf("%v\n", man)

	fmt.Println("Woman's Preference Matrix")
	fmt.Printf("%v\n", woman)

	StableMarriage(man, woman, n, mEng, nEng)

	fmt.Println("Man-Oriented Stable Marriage Engagement List")
	fmt.Println("Man  Woman     Woman  Man")
	fmt.Println("----------     ----------")
	for i := 0; i < n; i++ {
		fmt.Printf("%d%d%d%d\n", i, mEng[i], i, nEng[i])
	}
}