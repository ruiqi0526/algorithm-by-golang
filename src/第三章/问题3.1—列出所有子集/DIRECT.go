package main

import "fmt"

func main() {
	var (
		digit [20]string
		i int
		j int
		n int
	)

	fmt.Println("Direct Generation of All Subsets of a Set")
	fmt.Println("=========================================")
	fmt.Print("Number of Elements in the Given Set --> ")
	fmt.Scanln(&n)

	for i = 0; i < n; i++ {
		digit[i] = "0"
	}

	fmt.Println("{}")
	for {
		for i = 0; i < n && digit[i] == "1"; i++ {
			digit[i] = "0"
		}
		if i == n {
			break
		}else {
			digit[i] = "1"
		}

		for i = 0; i < n && digit[i] == "0"; i++ {}

		fmt.Printf("{%d", i+1)
		for j = i + 1; j < n; j++ {
			if digit[j] == "1" {
				fmt.Printf(",%d", j+1)
			}
		}
		fmt.Printf("}\n")
	}
}