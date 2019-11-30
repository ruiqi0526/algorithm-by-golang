package main

import (
	"fmt"
)

func head_tail(x []int, n int) int {
	var (
		prefix int = x[0]   //前置和
		suffix int = x[n - 1]   //后置和
		prefix_idx int = 0 
		suffix_idx int = n - 1
		count int = 0
	)

	for suffix_idx >= 0 && prefix_idx <= n - 1{
		if prefix < suffix {
			prefix_idx++
			prefix += x[prefix_idx]
		}else if prefix > suffix {
			suffix_idx--
			suffix += x[suffix_idx]
		}else{
			count++
			prefix_idx++
			suffix_idx--
			if prefix_idx >= n && suffix_idx <= 0{
				break
			}else{
				prefix += x[prefix_idx]
				suffix += x[suffix_idx]
			}
		}
	}

	return count
}

func main() {
	x := []int{3, 6, 2, 1, 4, 5, 2}
	n := len(x)

	fmt.Println("Given Array:", x)
	fmt.Printf("There are %d equal prefix-suffix sum pairs.\n", head_tail(x, n))
}