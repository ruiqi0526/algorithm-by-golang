package main

import "fmt"

func main() {
	var (
		factors = make([]int, 100)
		exps = make([]int, 100)
		n, work, count, i, k int
	)

	SAVE_FACTOR := func(fact, exp int) {
		if exp > 0 {
			factors[count] = fact
			exps[count] = i
			count++
		}
	}

	fmt.Println("Factorization by Division Program")
    fmt.Println("=================================")
	fmt.Print("Input a positive integer --> ")
	fmt.Scanln(&n)
	work = n
	for i = 0; (work & 1) == 0 && work > 1; i++ {
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
		fmt.Printf("%d(%d)\n", factors[i], exps[i])
	}

}