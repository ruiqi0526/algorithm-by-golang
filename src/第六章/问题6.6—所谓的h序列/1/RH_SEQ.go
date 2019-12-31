package main

import "fmt"
import "strings"

const (
	YES = 1
	NO = 0
)

var cursor int

func HSequence(x []string) int {
	length := len(x)

	if HSeq(x) == YES {
		if cursor == length - 1 {
			return YES
		}
	}
	return NO
}

func HSeq(x []string) int {
	if cursor >= len(x) {
		cursor = len(x) - 1
		return YES
	}
	
	switch x[cursor] {
		case "0":
			return YES
		case "1":
			cursor++
			if HSeq(x) == YES {
				cursor++
				if HSeq(x) == YES {
					return YES
				}
			}
			return NO
		default:
			return NO
	}
}

func main() {
	var line string
	fmt.Println("H-Sequence Testing Program (Recursive Version)")
	fmt.Println("==============================================")
	
	for done := NO; done == 0; {
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