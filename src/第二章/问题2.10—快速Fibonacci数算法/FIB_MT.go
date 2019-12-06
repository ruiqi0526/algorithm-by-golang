package main

import "fmt"

func matrix_power(a, b, c, d uint64, n int, aa, bb, cc, dd *uint64) {
	var xa, xb, xc, xd uint64

	if n == 1 {
		*aa = a
		*bb = b
		*cc = c 
		*dd = d
	}else if n & 1 == 1 {
		matrix_power(a, b, c, d, n - 1, &xa, &xb, &xc, &xd)
		*aa = a * xa + b * xc
		*bb = a * xb + b * xd
		*cc = c * xa + d * xc
		*dd = c * xb + d * xd
	}else{
		matrix_power(a, b, c, d, n >> 1, &xa, &xb, &xc, &xd)
		*aa = xa * xa + xb * xc
		*bb = xa * xb + xb * xd
		*cc = xc * xa + xd * xc
		*dd = xc * xb + xd * xd
	}
}

func fibonacci(n int) uint64 {
	var a, b, c, d uint64

	if n <= 2 {
		return 1
	}else{
		matrix_power(1, 1, 1, 0, n - 2, &a, &b, &c, &d)
		return a + b
	}
}


func main(){
	var n int

	fmt.Println("O(log(n)) Fibonacci Number Computation");
	fmt.Print("n please ---> ")
	fmt.Scanln(&n)
	fmt.Printf("fib(%d) = %d\n", n, fibonacci(n))
}