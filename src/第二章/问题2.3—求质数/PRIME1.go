package main

import "fmt"

const (
	MAXSIZE = 100
	YES = 1
	NO = 0
)

func main() {
	var (
		i, is_prime int
		gap int = 2	
		count int = 2
		may_be_prime int = 5
		prime [MAXSIZE]int
	)

	prime[0] = 2
	prime[1] = 3
	prime[2] = 5

	for count < MAXSIZE {
		may_be_prime += gap
		gap = 6 - gap
		is_prime = YES

		for i = 2; prime[i] * prime[i] <= may_be_prime && is_prime == 1; i++ {
			if may_be_prime % prime[i] == NO {
				is_prime = NO
			}
		}

		if is_prime == YES {
			count++
			if count < MAXSIZE {
				prime[count] = may_be_prime
			}else{
				break
			}
		}	
	}

	fmt.Printf("First %d prime numbers are :", count)
	for i = 0; i < count; i++ {
		if i % 10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%d\n", prime[i])
	}
}