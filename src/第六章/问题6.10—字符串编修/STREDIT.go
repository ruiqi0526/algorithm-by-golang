package main

import (
	"fmt"
	"strings"
)

const (
	INSERT_COST = 1
	DELETE_COST = 1
	EXCHANGE_COST = 2
)

func Min(x, y, z int) int {
	if x <= y {
		func() {
			if x <= z {
				return x
			} else {
				return z
			}
		}()
	} else {
		func() {
			if y <= z {
				return y
			} else {
				return z
			}
		}()
	}
}

func Edit(source, target []string, s, t []int, count *int) {
	var (
		sLen = len(source)
		tLen = len(target)
		insertT, deleteS, exchange int
		i, j, no int
		cost = make([][]int, 8 * (sLen + 1))
	)

	for i = 0; i < 8 * (sLen + 1); i++ {
		cost[i] = make([]int, 8 * (sLen + 1) * (sLen + 1))
	}

	for i = 1; i <= sLen; i++ {
		cost[i] = cost[i + tLen]
	}

	cost[0][0] = 0
	for i = 1; i <= sLen; i++ {
		cost[i][0] = cost[i - 1][0] + 1
	}
	for j = 1; j <= tLen; j++ {
		cost[0][j] = cost[0][j - 1] + 1
	}

	for i = 0; i < sLen; i++ {
		for j = 0; j < tLen; j++ {
			if source[i] == target[j] {
				cost[i + 1][j + 1] = cost[i][j]
			} else {
				insertT = cost[i + 1][j] + INSERT_COST
				deleteS = cost[i][j + 1] + DELETE_COST
				exchange = cost[i][j] + EXCHANGE_COST
				cost[i + 1][j + 1] = Min(insertT, deleteS, exchange)
			}
		}
	}

	i = sLen
	j = tLen
	for i != 0 && j != 0 {
		if cost[i][j] == cost[i - 1][j] + INSERT_COST {
			i--
		} else if cost[i][j] == cost[i][j - 1] + DELETE_COST {
			j--
		} else {
			s[no] = i - 1
			t[no] = j - 1
			no++
			i--
			j--
		}
	}
	Reverse(s, no)
	Reverse(t, no)
	*count = no
}

func Max(x, y int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func EditOp(source, target []string, s, t []int, n int) {
	var (
		length = Max(len(source), len(target)) + 1
		i, j, k int
		work = make([]string, length)
	)

	fmt.Println("Editing Operations are listed as Follows")
	for ; i < n; i++ {
		if s[i] - j > 0 {
			work = source[j : s[i] + 1]
			fmt.Printf("Delete \"%s\" from source string.\n", strings.Split(work, ""))
		}
		if t[i] - k > 0 {
			work = target[k : t[i] + 1]
			fmt.Printf("Insert \"%s\" to source string.\n", strings.Split(work, ""))
		}
		fmt.Printf("Keep \'%s\' from source string.\n", source[s[i]])
	}
	if len(source) - j > 0 {
		fmt.Printf("Delete \"%s\" from source string.\n", source[j])
	}
	if len(target) - k > 0 {
		fmt.Printf("Insert \"%s\" to source string.\n", target[k])
	}
	fmt.Println("DONE! Target string generated.\n")
}

func SWAP(a, b int) {
	a, b = b, a
}

func Reverse(x []int, n int) {
	i := 0
	j := n - 1
	for i <= j {
		SWAP(x[i], x[j])
		i++
		j--
	}
}

func main() {
	var (
		source string
		target string
		onSource = make([]int, 100)
		onTarget = make([]int, 100)
		equalCount int
	)

	fmt.Println("String Editing Problem")
	fmt.Println("======================")
	fmt.Print("Source String --> ")
	fmt.Scanln(&source)
	sourceArray := strings.Split(source, "")
	fmt.Print("Target String --> ")
	fmt.Scanln(&target)
	targetArray := strings.Split(target, "")
	Edit(sourceArray, targetArray, onSource, onTarget, &equalCount)
	EditOp(sourceArray, targetArray, onSource, onTarget, equalCount)
}