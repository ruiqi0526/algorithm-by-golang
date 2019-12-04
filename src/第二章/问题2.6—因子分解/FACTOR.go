package main

import "fmt"

var (
	factors [100]int
	exps [100]int
	n int
	work int
	count int
	i int
	k int
)

var SAVE_FACTOR = func(fact int, exp int) {
	if exp > 0 {
		factors[count] = fact
		exps[count] = exp
		count++
	}
}

func main() {
	fmt.Print("Input a positive integer --> ")
	fmt.Scanln(&n)

	work = n
	for i = 0; work > 1 && (work & 0x01) == 0; i++ {
		work >>= 1
	}

	SAVE_FACTOR(2, i)

	for k = 3; k <= work; k += 2 {
		for i = 0; work % k == 0 && work > 1; i++ {
			work /= k
		}
		SAVE_FACTOR(k, i)
	}

	fmt.Printf("%d = ", n)

	for i = 0; i < count; i++ {
		fmt.Printf("%d(%d)", factors[i], exps[i])
	}
}