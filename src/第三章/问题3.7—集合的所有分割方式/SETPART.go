package main

import "fmt"

func display(code []int, n int) {
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%d", code[i])
	}
}

func set_partition(n, sizeOfArray int) {
	var (
		code = make([]int, sizeOfArray)
		maxi = make([]int, sizeOfArray)
		i int
	)

	for i = 0; i < n; i++ {
		code[i] = 1
		maxi[i] = 2
	}

	for {
		display(code, n)
		for i = n - 1; code[i] == maxi[i]; i-- {}

		if i > 0 {
			code[i]++
			for j := i + 1; j < n; j++ {
				code[j] = 1
				maxi[j] = maxi[i] + func()(int){
					if code[i] == maxi[i] {
						return 1
					} else {
						return 0
					}
				}()
			}
		} else {
			break
		}
	}
}

func main(){	
	var n int

	fmt.Println("Set Partition Program for {1,2,3,...,N}")
	fmt.Println("=======================================")
	fmt.Print("N Please --> ")
	fmt.Scanln(&n)
	sizeOfArray := n * 8
	fmt.Println("Codes of Partitions.")
	fmt.Println("i-th position = j means i in partition j")
	set_partition(n, sizeOfArray)
}