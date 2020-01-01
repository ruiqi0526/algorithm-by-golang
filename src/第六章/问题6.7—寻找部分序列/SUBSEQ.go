package main

import (
	"fmt"
	"strings"
)

const (
	FOUND = 1
	NOT_FOUND = 0
)

func SubSequence(text, pat []string, loc []int) int {
	var (
		tLen = len(text)
		pLen = len(pat)
	)

	if pLen > tLen {
		return NOT_FOUND
	}
	i := 0
	for j := 0; j < tLen && j < pLen; j++ {
		for ; i < tLen && text[i] != pat[j]; i++ {}
		if i >= tLen {
			return NOT_FOUND
		} else {
			loc[j] = i
		}
	}
	return FOUND
}

func main() {
	var (
		t, p string
		loc = make([]int, 100)
		i, j int
	)

	fmt.Println("Subsequence Search Program")
	fmt.Println("==========================")
	fmt.Println("Text String    : ")
	fmt.Scanln(&t)
	tArray := strings.Split(t, "")
	fmt.Println("Pattern String : ")
	fmt.Scanln(&p)
	pArray := strings.Split(p, "")

	if SubSequence(tArray, pArray, loc) == FOUND {
		fmt.Println("Pattern's Locations are shown Bellow :")
		fmt.Printf("%s\n", t)
		for i = 0; i < len(pArray); i++ {
			j++
			for ; j < loc[i]; j++ {
				fmt.Print(" ")
			fmt.Printf("%s", tArray[i])
			}
		}
	} else {
		fmt.Println("*** Pattern NOT FOUND *****")
	}
}