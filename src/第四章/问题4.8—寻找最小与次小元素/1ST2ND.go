package main

import (
	"fmt"
	"math"
	"math/rand"
)

func first_second(x []int, n, first, second int) {
	f_and_s(x, 0, n - 1, first, second)
}

func f_and_s(x []int, left, right, f, s int) {
	var (
		mid int
		F1, F2 int
		S1, S2 int
	)

	if left > right {
		s = math.MaxInt64
		f = s
	} else if left == right {
		f = x[left]
		s = math.MinInt64
	} else {
		mid = (left + right) / 2
		f_and_s(x, left, mid, F1, S1)
		f_and_s(x, mid + 1, right, F2, S2)
		if F1 < F2 {
			f = F1
			s = func() int {
				if S1 < F2 {
					return S1
				} else {
					return F2
				}
			}()
		} else {
			f = F2
			s = func() int {
				if S1 < F2 {
					return S1
				} else {
					return F2
				}
			}()
		}
	}
}

func main(){
	var (
		x = make([]int, 100)
		n, first, second int
	)

	fmt.Println("Recursive First-Second Elements")
	fmt.Println("===============================")
	fmt.Print("How many elements (at least 2) --> ")
	fmt.Scanln(&n)

	fmt.Println("Randomly Generated Data :")
	for i := 0; i < n; i++ {
		x[i] = rand.Intn(100)
		if i % 100 == 0 {
			fmt.Println()
		}
		fmt.Printf("%d", x[i])
	}

	first_second(x, n, first, second)
	fmt.Printf("The Smallest Element --------->%6d\n", first);
    fmt.Printf("The Second Smallest Element -->%6d\n", second)
}