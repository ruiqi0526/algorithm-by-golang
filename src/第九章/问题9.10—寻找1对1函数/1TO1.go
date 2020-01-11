package main

import (
	"fmt"
)

const (
	QUEUE_SIZE = 100
	SAVED = 1
	DELETED = -1
)

func FindOneToOne(funct, status, counter []int, n int) {
	var (
		queue = make([]int, QUEUE_SIZE)
		head, tail int
	)

	for i := 0; i < n; i++ {
		status[i] = SAVED
	}

	for i := 0; i < n; i++ {
		counter[funct[i]]++
	}

	tail = -1
	for i := 0; i < n; i++ {
		if counter[i] == 0 {
			tail++
			queue[tail] = i
		}
	}
	head++
	for head <= tail {
		j := queue[head]
		head++
		status[j] = DELETED
		counter[funct[j]]--
		if counter[funct[j]] == 0 {
			tail++
			queue[tail] = funct[j]
		}
	}
}

func main() {
	var (
		functTable = []int{2, 0, 0, 4, 4, 3, 5}
		n = len(functTable) / 4
		status = make([]int, n)
		counter = make([]int, n)
	)

	fmt.Println("One-To-One Function Construction Program")
	fmt.Println("========================================")
	fmt.Println("Domain    Range    Status")
	fmt.Println("------    -----    ------")

	FindOneToOne(functTable, status, counter, n)

	for i := 0; i < n; i++ {
		fmt.Printf("%4d%10d\n", i, functTable[i])
		if status[i] == SAVED {
			fmt.Println("        SAVED")
		} else {
			fmt.Println("       DELETED")
		}
	}

	fmt.Println("Constructed New 1-1 Function")
	fmt.Println("Domain    Range")
	fmt.Println("------    -----")
	for i := 0; i < n; i++ {
		if status[i] == SAVED {
			fmt.Printf("%4d%10d\n", i, functTable[i])
		}
	}
}