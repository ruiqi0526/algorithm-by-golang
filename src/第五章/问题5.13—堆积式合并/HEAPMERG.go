package main

import (
	"fmt"
	"math"
)

const (
	YES = 1
	NO = 0
)

var (
	file_no int
	file_len int
	heaq []int
	location []int
	index []int
)

func heap_merge(matrix [][]int, m, n int, out []int) {
	var row int

	prepare(matrix, m, n)
	out_no := 0
	temp := delete_top(&row)
	for temp != math.MaxInt64 {
		out_no++
		out[out_no] = temp
		insert(ask_data(matrix, row), row)
	}
}

func prepare(matrix [][]int, m, n int) {
	heaq = make([]int, 8 * (m + 1))
	location = make([]int, 8 * (m + 1))
	index = make([]int, m)
	file_no = m
	file_len = n

	for i := 1; i <= file_no; i++ {
		heaq[i] = matrix[i - 1][0]
		location[i] = i - 1
		index[i - 1] = 1
	}
}

func ask_data(matrix [][]int, row int) int {
	if index[row] >= file_len {
		return math.MaxInt64
	} else {
		index[row]++
		return matrix[row][index[row]]
	}
}

func delete_top(row *int) int {
	*row = location[1]
	return heaq[1]
}

func insert(data, row int) {
	location[1] = row
	heaq[1] = data
	siftdown()
}

func makeheaq() {
	for i := file_no / 2; i >= 1; i-- {
		fix_heaq(i)
	}
}

func siftdown() {
	fix_heaq(1)
}

func fix_heaq(root int) {
	key_heaq := heaq[root]
	key_loc := location[root]
	father := root
	son := father + father
	done := false

	for son <= file_no && !done {
		if son < file_no && heaq[son + 1] < heaq[son] {
			son++
		} 
		if key_heaq > heaq[son] {
			heaq[father] = heaq[son]
			location[father] = location[son]
			father = son
		} else {
			done = true
		}
		son = father + father
	}
	heaq[father] = key_heaq
	location[father] = key_loc
}

func main() {
	var matrix = [][]int{
		{  1,  3,  5,  7,  9, 10, 15, 20, 50, 80},
		{  5,  6,  9, 11, 14, 19, 23, 45, 67, 78},
		{  2, 10, 20, 23, 39, 47, 60, 75, 82, 90},
		{  9, 18, 27, 42, 55, 76, 83, 87, 88, 91},
		{  4,  9, 19, 29, 39, 49, 59, 69, 79, 99},
	}
	var m int = 5
	var n int = 10
	var out = make([]int, 20 * 20)

	fmt.Println("Heap Merge Program")
	fmt.Println("==================")
	
	for i := 0; i < m; i++ {
		fmt.Printf("\nFile #%d :", i)
		for j := 0; j < n; j++ {
			if j % 10 == 0 {
				fmt.Println()
			}
			fmt.Printf("%d ", matrix[i][j])
		}
	}

	heap_merge(matrix, m, n, out)
	fmt.Println("The Merged Result :")
	for i := 0; i < n * m; i++ {
		if i % 10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%d ", out[i])
	}
}