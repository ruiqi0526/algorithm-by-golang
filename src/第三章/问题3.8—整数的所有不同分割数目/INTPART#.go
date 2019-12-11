package main

import "fmt"

func int_part_no(n int) uint64 {
	return compute(n, n)
}

func compute(number int, maximum int) uint64 {
	if number == 1 || maximum == 1 {
		return 1
	}else if number < maximum {
		return compute(number, number)
	}else if number == maximum {
		return 1 + compute(number, maximum - 1)
	}else {
		return compute(number, maximum - 1) + compute(number - maximum, maximum)
	}
}

func main() {
	var n int

	fmt.Println("Number of partitions of an Integer")
	fmt.Println("==================================")
	fmt.Print("N --> ")
	fmt.Scanln(&n)
	fmt.Printf("There are %d partitions.\n", int_part_no(n))
}