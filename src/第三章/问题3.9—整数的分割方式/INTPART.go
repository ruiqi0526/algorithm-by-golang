package main

import "fmt"

func display(partition [20]int, mult [20]int, part_no int) {
	for i := 0; i <= part_no; i++ {
		for j := 0; j < mult[i]; j++ {
			fmt.Printf("%d\n", partition[i])
		}
	}
}

func main() {
	var (
		partition [20]int
		mult [20]int
		part_no int
		sum, size, remainder, count int
		n int
	)

	fmt.Println("Partition of Integer")
	fmt.Println("====================")
	fmt.Print("Input a Positive Integer --> ")
	fmt.Scanln(&n)

	partition[1] = n
	mult[1] = 1
	part_no = 1
	count = 1
	display(partition, mult, part_no)

	for {
		sum = func() int {
			if partition[part_no] == 1 {
				part_no--
				return mult[part_no] + 1
			}else {
				return 1
			}
		}()

		size = partition[part_no] - 1
		if mult[part_no] != 1 {
			part_no++
			mult[part_no]--
		}
		partition[part_no] = size
		mult[part_no] = sum / size + 1
		remainder = sum % size
		if remainder != 0 {
			part_no++
			partition[part_no] = remainder
			mult[part_no] = 1
		}
		count++
		display(partition, mult, part_no)

		if mult[part_no] != n {
			fmt.Printf("There are %d partitions in anti-lexical order", count)
		} else {
			break
		}
	}
}