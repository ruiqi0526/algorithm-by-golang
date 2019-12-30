package main

import (
	"fmt"
	"strings"
)

func Setup(pat []string, fail []int) {
	var length int = len(pat)
	var i, j int

	fail[0] = -1
	for i = 1; i < length; i++ {
		for j = fail[i - 1]; j >= 0 && pat[j + 1] != pat[i]; j = fail[j] {}
		fail[i] = func() int {
			if j < 0 && pat[j + 1] != pat[i] {
				return -1
			} else {
				return j + 1
			}
		}()
	}
}

func KMP(text, pat []string, fail []int) int {
	var (
		tLength int = len(text)
		pLength int = len(pat)
		t, p int
	)

	Setup(pat, fail)
	for t = 0; t < tLength && p < pLength; {
		if text[t] != pat[p] {
			if p > 0 {
				p = fail[p - 1] + 1
			} else {
				t++
			}
		} else {
			t++
			p++
		}
	}

	res := func() int {
		if p >= pLength {
			return t - pLength
		} else {
			return -1
		}
	}()

	return res
}

func main() {
	var (
		text string
		pat string
		fail = make([]int, 100)
	)
	
	fmt.Println("Knuth-Morris-Pratt String Search Program")
	fmt.Println("========================================")
	fmt.Println("Text String ------> ")
	fmt.Scanln(&text)
	textArray := strings.Split(text, "")
	fmt.Println("Pattern String ---> ")
	fmt.Scanln(&pat)
	patArray := strings.Split(pat, "")

	answer := KMP(textArray, patArray, fail) 
	if answer >= 0 {
		fmt.Println()
		for i := 1; i <= 6; i++ {
			fmt.Printf(" %d", i)
		}

		fmt.Println()
		for i := 1; i <= 6; i++ {
			fmt.Print("0....5....")
		}
		fmt.Println("0")
		fmt.Printf("%s\n", strings.Join(textArray, ""))
		for i := 0; i < answer; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("%s\n", strings.Join(patArray, ""))
		fmt.Printf("Pattern Found at location %d\n", answer)
	} else {
		fmt.Println("Pattern NOT found")
	}
}