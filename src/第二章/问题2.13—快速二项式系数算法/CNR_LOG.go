package main

import "fmt"

func power(m int, n uint64) int {
	temp := 1

	for n > 0 {
		if n & 1 == 1 {
			temp *= m
		}
		m *= m
		n >>= 1
	}

	return temp
}

func cnr(n uint64, answer []int) {
	var i uint64
	x := (1 << n) + 1
	mask := (1 << n) - 1
	
	result := power(x, n)

	for i = 0; i <= n; i++ {
		result <<= n
		answer[i] = result & mask
	}
}

func main() {
	var (
		n uint64
		r uint64
		answer = make([]int, 10)
	)

	fmt.Println("Fast Combinatorial Coefficient Computation")
    fmt.Println("==========================================")
	fmt.Print("N ---> ")
	fmt.Scanln(&n)

	cnr(n, answer)
	fmt.Println("All Combinatorial Coefficients :")
	for r = 0; r <= n; r++ {
		fmt.Printf("C(%d, %d) = %d\n", n, r, answer[r])
	}
}