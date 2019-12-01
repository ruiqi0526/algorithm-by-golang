package main

import "fmt"

func main() {
	var (
		p int
		q int
		r int
		count int
	)

	for number := 100; number <= 999; number++ {
		p = number / 100
		q = (number % 100) / 10
		r = number % 10
		if p * p * p + q * q * q + r * r * r == number {
			count++
			fmt.Printf("%d : %d\n", count, number)
		}
	}

	fmt.Printf("There are %d 3-digit Armstrong Numbers.\n", count)
}