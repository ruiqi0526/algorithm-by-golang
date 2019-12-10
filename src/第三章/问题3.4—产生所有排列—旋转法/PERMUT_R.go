package main

import "fmt"

var (
	perm [20]int
	position int
	n int
)

func ROTATE(p int) {
	temp := perm[p]
	for i := p - 1; i >= 0; i-- {
		perm[i + 1] = perm[i]
	}
	perm[0] = temp
}

func main() {
	fmt.Println("Permutation by Rotation Method");
	fmt.Println("==============================");
	fmt.Print("Number of Elements --> ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		perm[i] = i + 1
	}

	position = n - 1
	for position != 0 {
		fmt.Println()
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", perm[i])
		}
		position = n - 1
		ROTATE(position)
		for perm[position] == position + 1 && position != 0 {
			position--
			ROTATE(position)
		}
	}
}