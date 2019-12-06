package main

import "fmt"

func power(m, n uint64) uint64 {
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

func cn2(n uint64) uint64 {
	var x uint64 = (1 << n) + 1
	var mask uint64 = (1 << n) - 1

	return (power(x, n) >> ((n >> 1) * n)) & mask
}

func factorial(n uint64) uint64 {
	if n == 1 {
		return 1
	} else if n & 1 == 1 {
		return n * factorial(n - 1)
	} else {
		temp := factorial(n >> 1)
		return cn2(n) * temp *temp
	}
}

func main() {
	var n uint64

	fmt.Println("Fast (but NOT THE FASTEST) N! Computation :");
	fmt.Print("N for N! please --> ")
	fmt.Scanln(&n)
	fmt.Printf("%d! = %d\n", n, factorial(n))
}