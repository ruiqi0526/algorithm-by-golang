package main

import "fmt"

func cnr(n, r int) int {
	var c [100]int

	for i := 0; i <= r; i++ {
		c[i] = 1
	}

	for i := 0; i <= n - r; i++ {
		for j := 1; j <= r; j++ {
			c[j] += c[j-1]
		}
	}
	return c[r]
}

func main() {
	var n, r int

	fmt.Println("Combinatorial Coefficient by Addition")
    fmt.Println("=====================================");
	fmt.Print("N ---> ")
	fmt.Scanln(&n)
	fmt.Print("R ---> ")
	fmt.Scanln(&r)

	fmt.Printf("C(%d, %d) = %d\n", n, r, cnr(n, r))
}