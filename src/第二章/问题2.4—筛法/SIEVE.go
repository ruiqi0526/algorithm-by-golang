package main

import "fmt"

const (
	MAXSIZE = 200
	DELETED = 1
	KEPT = 0
)

func main() {
	var (
		sieve [MAXSIZE + 1]int
		count int = 0
		prime int
		k int
	)

	fmt.Printf("Prime Numbers between 2 and %d\n", MAXSIZE * 2 + 3)

	for i := 0; i <= MAXSIZE; i++ {
		sieve[i] = KEPT
	}

	for i, v := range sieve {
		if v == KEPT {
			prime = i + i + 3
			count++
			for k = prime + i; k <= MAXSIZE; k += prime {
				sieve[k] = DELETED
			}
		}
	}

	fmt.Println(2)
	k = 2
	for i, v := range sieve {
		if v == KEPT {
			if k > 10 {
				k = 1
			}
			fmt.Printf("%d\n", 2 * i + 3)
			k++
		}
	}

	fmt.Printf("There are %d primes in total.\n", count)
}