package main

import "fmt"

func fibonacci(n int) uint64 {
	var f0, f1, temp uint64
	var i int

	if n <= 2 {
		return 1
	}else{
		f0 = 1
		f1 = 1
		for i = 3; i <= n; i++ {
			temp = f0 + f1
			f0 = f1
			f1 = temp
		}
		return f1
	}
}

func main() {
	var n int

	fmt.Println("Iterative Linear Fibonacci Number Computation");
	fmt.Print("n please ---> ")
	fmt.Scanln(&n)
	fmt.Printf("fib(%d) = %v\n", n, fibonacci(n))
}