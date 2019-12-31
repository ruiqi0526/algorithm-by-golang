package main

import (
	"fmt"
	"strings"
)

const (
	YES = 1
	NO = 0
)

func HSequence(x []string) int {
	var (
		length int = len(x)
		count int = 1
		i int
	)

	for i = 0; count != 0 && i < length; i++ {
		switch x[i] {
		case "0":
			count--
		case "1":
			count++
		default:
			return NO
		}
	}
	if count == 0 && i >= length {
		return YES
	} else {
		return NO
	}
}

func main() {
	var (
		line string
		done int
	)

	fmt.Println("H-Sequence Testing Program (Iterative Version)")
	fmt.Println("==============================================")
	for done == NO {
		fmt.Print("Input a string of 0 and 1 --> ")
		fmt.Scanln(&line)
		if line != "" {
			lineArray := strings.Split(line, "")
			if HSequence(lineArray) == YES {
				fmt.Println("*** Input is a H sequence ***")
			} else {
				fmt.Println("*** Input is NOT a H sequence ***")
			}
		} else {
			done = YES
		}
	} 
}