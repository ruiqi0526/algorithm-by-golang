package main

import (
	"fmt"
	"math"
	"math/rand"
)

var (
	number int
	in1_data []int
	in2_data []int
	out_data []int
	in1_count []int
	in2_count []int
	out_count []int
)

func MIN(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func SWAP(a, b interface{}) {
	a, b = b, a
}

func LOG_sort(x []int, n int) {
	var log_n = int(math.Log(float64(n)) + 0.5)

	in1_data  = make([]int, 8 * log_n)
	in2_data  = make([]int, 8 * log_n)
	out_data  = make([]int, 8 * log_n)
	in1_count = make([]int, 8 * log_n)
	in2_count = make([]int, 8 * log_n)
	out_count = make([]int, 8 * log_n)

	number = LOG_merge(x, n, log_n)
	LOG_expand(x)
}

func LOG_merge(x []int, n, seg_size int) int {
	var (
		no1, no2 int
		start, len int
		i, j, k int
	)

	no1 = compress(x, seg_size, in1_data, in1_count)

	for start = seg_size; start < n; start += len {
		len = MIN(seg_size, n - start)
		no2 = compress(x, len, in2_data, in2_count)
		j = 0; k = 0
		for i = 0; i < no1 && j < no2; {
			if in1_data[i] < in2_data[j] {
				out_data[k] = in1_data[i]
				out_count[k] = in1_count[i]
				i++; k++
			} else if in1_data[i] > in2_data[j] {
				out_data[k] = in2_data[j]
				out_count[k] = in2_count[j]
				j++; k++
			} else {
				out_data[k] = in1_data[i]
				out_count[k] = in1_count[i] + in2_count[j]
				i++; j++; k++
			}
		}
		for ; i < no1; i++ {
			out_data[k] = in1_data[i]
			out_count[k] = in1_count[i]
			k++
		}
		for ; j < no2; j++ {
			out_data[k] = in2_data[j]
			out_count[k] = in2_count[j]
			k++
		}
		no1 = k
		SWAP(in1_data, out_data)
		SWAP(in1_count, out_count)
	}
	return no1
}

func compress(x []int, n int, data []int, count []int) int {
	var no = 0

	if n == 0 {
		return 0
	} else {
		data[0] = x[0]
		count[0] = 1
		for i := 1; i < n; i++ {
			if x[i] == x[i - 1] {
				count[no]++
			} else {
				no++
				data[no] = x[i]
				count[no] = 1
			}
		}
		no++
		return no
	}
}

func LOG_expand(x []int) {
	total := 0
	for i := 0; i < number; i++ {
		for j := 0; j < in1_count[i]; j++ {
			total++
			x[total] = in1_data[i]
		}
	}
}

func main() {
	var (
		x = make([]int, 100)
		n int
		temp int
	)

	fmt.Println("Array Sorting with Duplicated Distinct Elements")
	fmt.Println("===============================================")
	fmt.Println("How many elements do you want ? ")
	fmt.Scanln(&n)
	log_n := int(math.Log(float64(n)) + 0.5)

	for i := 0; i < n; {
		temp = rand.Intn(100)
		for j := 0; j < n / log_n && i < n; j++ {
			i++
			x[i] = temp
		}
	}

	for i := 0; i < n; i++ {
		j := rand.Intn(100)
		x[i], x[j] = x[j], x[i]
	}

	fmt.Println("Generated Array :")
	for i := 0; i < n; i++ {
		if i % 10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%d ", x[i])
	}

	LOG_sort(x, n)
	fmt.Println("Sorted Array :")
	for i := 0; i < n; i++ {
		if i % 10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%d ", x[i])
	}
}