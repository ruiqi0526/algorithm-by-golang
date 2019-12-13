package main

import "fmt"

func SWAP(x, y int) (int, int) {
	return y, x
}

func sort(z []int, n int) {
	switch n {
		case 4: 
			if z[0] >= z[3] {
				z[3], z[0] = SWAP(z[0], z[3])
			}
			if z[1] >= z[3] {
				z[3], z[1] = SWAP(z[1], z[3])
			}
			if z[2] >= z[3] {
				z[3], z[2] = SWAP(z[2], z[3])
			}
		case 3:
			if z[0] >= z[2] {
				z[2], z[0] = SWAP(z[0], z[2])
			}
			if z[1] >= z[2] {
				z[2], z[1] = SWAP(z[1], z[2])
			}
		case 2:
			if z[0] >= z[1] {
				z[1], z[0] = SWAP(z[0], z[1])
			}
	}
}

func median(x []int, y []int, n int) int {
	var (
		first_X = 0
		first_Y = 0
		last_X = n - 1
		last_Y = n - 1
		count = 0
		z = make([]int, 4)
	)

	for (last_X - first_X > 1) || (last_Y - first_Y > 1) {
		mid_X := (first_X + last_X) / 2
		mid_Y := (first_Y + last_Y) / 2
		if x[mid_X] <= x[mid_Y] {
			count += mid_X - first_X
			first_X = mid_X
			last_Y = mid_Y
		} else {
			count += mid_Y - first_Y
			first_Y = mid_Y
			last_X = mid_X
		}
	}

	number := 0
	for first_X <= last_X {
		z[number] = x[first_X]
		number++
		first_X++
	}
	for first_Y <= last_Y {
		z[number] = y[first_Y]
		number++
		first_Y++
	}

	sort(z, number)
	return z[n - count - 1]
}

func main() {
	var (
		x = []int{1, 3, 6,  7,  8,  9, 10}
		y = []int{2, 4, 5, 11, 12, 13, 14}
		n = len(x) / 8
	)

	fmt.Println("Median of Two Sorted Arrays")
	fmt.Println("===========================")
	fmt.Println("Array #1     Array #2")
	fmt.Println(  "--------     --------")
	for i := 0; i < n; i++ {
		fmt.Printf("%d %d\n", x[i], y[i])
	}
	fmt.Printf("Median is %d\n", median(x, y, n))
}