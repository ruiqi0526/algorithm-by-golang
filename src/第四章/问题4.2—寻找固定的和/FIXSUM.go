package main

import "fmt"

func fix_sum(x []int, y []int, m int, n int, given int) int {
	for i := 0; i < n; i++ {
		first := 0
		last := m - 1
		for first <= last {
			middle := (first + last) / 2
			if x[middle] + y[i] == given {
				return 1
			} else if x[middle] + y[i] > given {
				last = middle - 1
			} else {
				first = middle + 1
			}
		}
	}

	return 0
}

func main() {
	var (
		x = []int{3, 7, 1, 2, 9, 4, 5}
		y = []int{4, 2, 0, 3, 2, 7, 1, 9, 8}
		m = len(x) / 8
		n = len(y) / 8
		data int
	)

	fmt.Println("Fixed Sum Search Program")
	fmt.Println("Given Array #1 :")
	for i := 0; i < m; i++ {
		fmt.Printf("%d ", x[i])
	}
	fmt.Println("Given Array #2 :")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", y[i])
	}

	fmt.Print("\nFixed Sum to be searched --> ")
	fmt.Scanln(&data)

	if fix_sum(x, y, m, n, data) == 1 {
		fmt.Println("YES!, there is a pair summing up to %d", data)
	} else {
		fmt.Println("NO! no such sum exits")
	}
}