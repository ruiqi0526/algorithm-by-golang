package main

import "fmt"

func main() {
	var (
		set [20]int
		n, position int
	)

	fmt.Println("All Possible Subsets Generation by Lexical Order");
	fmt.Println("================================================");
	fmt.Print("Number of Elements in the Set --> ")
	fmt.Scanln(&n)

	fmt.Println("{}")
	set[position] = 1

	for {
		fmt.Printf("{%d", set[0])
		for i := 0; i <= position; i++ {
			fmt.Printf(",%d", set[i])
		}
		fmt.Println("}")

		if set[position] < n {
			set[position + 1] = set[position] + 1
			position++
		}else if position != 0 {
			position -= 1
			set[position]++
		}else {
			break
		}
	}
}