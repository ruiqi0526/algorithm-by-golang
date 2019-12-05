package main

import "fmt"

func iterative_power(m, n uint64) uint64 {
	var temp uint64 = 1

	for n > 0 {
		if n & 1 == 1 {
			temp *= m
		}
		m *= m
		n >>= 1
	}

	return temp
}


func main() {
	var m, n uint64

	fmt.Println("M ^ N Computation (M > 0 and N > 0)");
	fmt.Print("M = ")
	fmt.Scanln(&m)
	fmt.Print("N = ")
	fmt.Scanln(&n)

	fmt.Printf("M ^ N = %v\n", iterative_power(m, n))
}