package main

import (
	"fmt"
	"strings"
)

const (
	MAXLENGTH = 100
	YES = 1
	NO = 0
)

func parCount(line string, error []string, sw *int) {
	var (
		location = make([]int, MAXLENGTH)
		left = 0
		right = 0
		locPtr = 1
		i int
	)

	*sw = NO
	lineSplit := strings.Split(line, "")
	for i = 0; i < len(lineSplit); i++ {
		error[i] = " "
		if lineSplit[i] == "(" {
			locPtr++
			location[locPtr] = i
			left++
		} else if lineSplit[i] == ")" {
			if left <= right {
				error[i] = "?"
				*sw = YES
			} else {
				right++
				locPtr--
			}
		}
	}

	if locPtr >= 0 {
		*sw = YES
		for i = 0; i <= locPtr; i++ {
			error[location[i]] = "$"
		}
	}
}

func main() {
	var (
		error = make([]string, MAXLENGTH)
		errorSw int 
		line string
	)

	fmt.Println("Parenthesis Counting Program")
	fmt.Println("============================")
	fmt.Println("Input a line please")
	fmt.Scanln(&line)

	parCount(line, error, &errorSw)
	if errorSw == NO {
		fmt.Println("Correct Input")
	} else {
		fmt.Printf("%s\n", strings.Join(error, ""))
	}
}