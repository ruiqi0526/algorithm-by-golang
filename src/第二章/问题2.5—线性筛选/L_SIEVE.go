package main

import "fmt"

const MAXSIZE = 1000

var (
	previous [MAXSIZE + 1]int
	next [MAXSIZE + 1]int
	prime int
	fact int
	i int
	mult int
	n int
	count int
)

var REMOVE = func(x int) {
	next[previous[x]] = next[x]
	previous[next[x]] = previous[x]
}

var INITIAL = func(x int) {
	var i int
	for i = 2; i <= n; i++ {
		previous[i] = i - 1
		next[i] = i + 1
	}
	next[n] = -1
	previous[2] = next[n]
}

func main() {
	fmt.Print("Prime Numbers between 2 to --> ")
	fmt.Scanln(&n)

	INITIAL(n)
	for prime = 2; prime * prime <= n; prime = next[prime] {
		for fact = prime; fact * prime <= n; fact = next[fact] {
			for mult = fact * prime; mult <= n; mult *= prime {
				REMOVE(mult)
			}
		}
	}

	for i = 2; i != -1; i = next[i] {
		if count % 8 == 0 {
			print("\n")
		}
		fmt.Printf("%d\n", i)
		count++
	}

	fmt.Printf("There are %ld Prime Numbers in Total\n", count)
}