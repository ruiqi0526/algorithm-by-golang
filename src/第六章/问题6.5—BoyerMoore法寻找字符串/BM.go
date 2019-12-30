package main

import (
	"fmt"
	"strings"
)

const NOT_FOUND = -1

func BM(text, pat []string) int {
	var (
		jumpTab = make([]int, 256)
		tLen = len(text)
		pLen = len(pat)
		i, j, k int
	)

	GetJump(pat, jumpTab)
	for i = pLen - 1; i < tLen; {
		k = i
		for j = pLen - 1; j >= 0 && text[k] == pat[j]; j-- {
			k--
		}
		if j < 0 {
			return k + 1
		} else {
			i += jumpTab[[]rune(text[i])[0]]
		}
	}
	return NOT_FOUND
}

func GetJump(pat []string, jumpTab []int) {
	var length int = len(pat)

	for i := 1; i < 256; i++ {
		jumpTab[i] = length
	}
	for i := 0; i < length - 1; i++ {
		jumpTab[[]rune(pat[i])[0]] = length - i - 1
	}
}

func main() {
	var (
		text string
		pat string
	)

	fmt.Println("Boyer-Moore String Searching Program")
	fmt.Println("====================================")
	fmt.Println("Text String    --> ")
	fmt.Scanln(&text)
	textArray := strings.Split(text, "")
	fmt.Println("Pattern String --> ")
	fmt.Scanln(&pat)
	patArray := strings.Split(pat, "")

	answer := BM(textArray, patArray)
	if answer >= 0 {
		fmt.Println()
		for i := 1; i <= 6; i++ {
			fmt.Printf(" %d", i)
		}
		fmt.Println()
		for i := 1; i <= 6; i++ {
			fmt.Println("0....5....")
		}
		fmt.Println("0")
		fmt.Printf("%s\n", strings.Join(textArray, ""))
		for i := 0; i < answer; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("%s\n", strings.Join(patArray, ""))
		fmt.Printf("Pattern Found at location %d\n", answer)
	} else {
		fmt.Println("Pattern NOT FOUND");

	}
}