package main

import (
	"fmt"
)

const (
	YES = 1
	NO = 0
)

func Sink(know [][]int, n int) int {
	var (
		candidate int
		wrong bool
		theSink int
		i, k int
	)

	theSink = -1
	for i = 1; i < n; i++ {
		if know[candidate][i] == 1 {
			candidate = i
		}
	}
	wrong = false
	for k = 0; k < n && !wrong; k++ {
		if know[candidate][k] == 1 {
			wrong = true
		}
		if know[candidate][k] == 0 && candidate != k {
			wrong = true
		}
	}
	if !wrong {
		theSink = candidate
	}
	return theSink
}

func main() {
	var know = [][]int{
		{ NO,  NO, YES, YES, YES},
		{YES,  NO, YES, YES, YES},
		{ NO, YES,  NO, YES,  NO},
		{ NO,  NO,  NO,  NO,  NO},
		{YES,  NO, YES, YES, YES},
	}
	var n = 5
	var answer int

	fmt.Println("Sink Finding Program")
    fmt.Println("====================")
	fmt.Println("Given know[*,*] Matrix:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf(" %s", func() string {
				if know[i][j] == 1 {
					return " Y"
				} else {
					return " N"
				}
			}())
		}
	}
	answer = Sink(know, n)
	if answer >= 0 {
		fmt.Printf("The sink is %d\n", answer + 1)
	} else {
		fmt.Println("NO!, there is no sink")
	}
}