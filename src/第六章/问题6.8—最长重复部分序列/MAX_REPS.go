package main

import (
	"fmt"
	"strings"
)

const (
	FOUND = 1
	NOT_FOUND = 0
)

func MaxRepetition(text, pat []string, loc []int) int {
	var (
		tLen = len(text)
		pLen = len(pat)
		low = 1
		high = tLen / pLen
	)

	if SubSequence(text, pat, low, loc) == NOT_FOUND {
		return NOT_FOUND
	}

	for low + 1 < high {
		mid := (low + high) / 2
		if SubSequence(text, pat, mid, loc) == FOUND {
			low = mid
		} else {
			high = mid
		}
	}

	if SubSequence(text, pat, high, loc) == FOUND {
		return high
	}
	return low
}

func SubSequence(text, pat []string, number int, loc []int) int {
	var (
		tLen = len(text)
		pLen = len(pat)
		count, index, i, j int
	)

	if pLen > tLen {
		return NOT_FOUND
	}

	for ; i < tLen && j < pLen; j++ {
		for count < number {
			for ; i < tLen && j < pLen; j++ {}
			if i >= tLen {
				return NOT_FOUND
			} else {
				count++
				i++
				loc[index] = i
			}
		}
	}
	return FOUND
}

func main() {
	var (
		t, p string
		loc = make([]int, 100)
		n, i, j int
	)

	fmt.Println("Maximum Repetition of Subsequence")
	fmt.Println("=================================")
	fmt.Println("Text String    : ")
	fmt.Scanln(&t)
	tArray := strings.Split(t, "")
	fmt.Println("Pattern String : ")
	fmt.Scanln(&p)
	pArray := strings.Split(p, "")

	n = MaxRepetition(tArray, pArray, loc)
	if n > 0 {
		fmt.Printf("The Maximum Repetition Factor is %d\n", n);
        fmt.Println("Pattern's Locations are shown Bellow :");
		fmt.Printf("%s\n", t)
		for ; i < len(pArray) * n; i++ {
			j++
			for ; j < loc[i]; j++ {
				fmt.Print(" ")
			fmt.Printf("%s", t[j])
			}
		}
	} else {
		fmt.Println("*** Pattern NOT FOUND *****")
	}
}