package main

import "fmt"

func ext_fibonacci(n, x, y int) int {
	var f0, f1, a0, a1, temp1, temp2 int
	
	if n == 0{
		return 1
	} else if n == 1 {
		return 2
	} else {
		f0 = 1
		f1 = 1
		a0 = 1
		a1 = 2
		for i := 2; i <= n; i++ {
			temp1 = x * f1 + y * f0
			f0    = f1
			f1    = temp1
			temp2 = x * a0 + y * (a1 - f0) + f0 + f1
			a0    = a1
			a1    = temp2
		}
		return a1
	}
}

func main(){
	var n, x, y int

	fmt.Print("x ---> ")
	fmt.Scanln(&x)
	fmt.Print("y ---> ")
	fmt.Scanln(&y)
	fmt.Print("n ---> ")
	fmt.Scanln(&n)
	fmt.Printf("Answer is %d\n", ext_fibonacci(n, x, y))
}