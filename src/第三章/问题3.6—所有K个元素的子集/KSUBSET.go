package main

import "fmt"

func main() {
	var (
		set [20]int
		n, k int
	)

	fmt.Println("All k Elements Subsets from a n Elements Set")
	fmt.Println("============================================")
	fmt.Print("Number of Elements in the Set --> ")
	fmt.Scanln(&n)
	fmt.Print("Number of Elements in Subset ---> ")
	fmt.Scanln(&k)

	for i := 0; i < k; i++ {
		set[i] = i + 1
	}
	fmt.Printf("{%d", set[0])

	for j := 1; j < k; j++ {
		fmt.Printf(",%d", set[j])
	}
	fmt.Println("}")

	position := k - 1
	for {
		if set[k - 1] == n{
			position--
		}else {
			position = k - 1
		}

		set[position]++
		for i := position + 1; i < k; i++ {
			set[i] = set[i - 1] + 1
		}

		fmt.Printf("{%d", set[0])
		for j := 1; j < k; j++ {
			fmt.Printf(",%d", set[j])
		}
		fmt.Println("}")

		if set[0] >= n - k + 1 {
			break
		}
	}

}