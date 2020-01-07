package main

import (
	"fmt"
)

var count int

func ParentheseGen(n int) int {
	var x = make([]string, n)

	count = 0
	Gen(x, n / 2, n / 2, 0)

	return count
}

func Gen(x []string, n, leftNo, rightNo int) {
	var i int

	if leftNo == 0 && rightNo == 0 {
		fmt.Printf("%d   %s\n", count, x)
		count++
	} else {
		i = 2 * (n - leftNo) - rightNo
		if rightNo != 0 {
			x[i] = ")"
			Gen(x, n, leftNo, rightNo - 1)
		}
		if leftNo != 0 {
			x[i] = "("
			Gen(x, n, leftNo - 1, rightNo + 1)
		}
	}
}

func main() {
	var (
		n int
		answer int
	)

	fmt.Println("Well-Formed Parenthesis Generation")
    fmt.Println("==================================")
	fmt.Print("String Length --> ")
	fmt.Scanln(&n)

	if n % 2 != 0 {
		fmt.Println("*** ERROR ***  String Length must be even.")
	} else {
		fmt.Println("   No   Generated String")
		fmt.Println("  ---   ----------------")
		answer = ParentheseGen(n)
		fmt.Printf("There are %d Parentheses Strings Generated.\n", answer)
	}
}