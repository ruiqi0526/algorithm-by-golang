package main

import "fmt"

func recursive_power(m, n uint64) uint64 {
	var temp uint64

	if n == 0 {
		return 1
	}else if n & 0x01 == 0 {
		temp = recursive_power(m, n >> 1)
		return temp * temp
	}else{
		return m * recursive_power(m, n - 1)
	}
}

func main() {
	var m, n uint64

	fmt.Println("M ^ N Computation (M > 0 and N > 0)");
	
	fmt.Print("M = ")
	fmt.Scanln(&m)
	fmt.Print("N = ")
	fmt.Scanln(&n)
	
	fmt.Printf("M ^ N = %v\n", recursive_power(m, n))
}