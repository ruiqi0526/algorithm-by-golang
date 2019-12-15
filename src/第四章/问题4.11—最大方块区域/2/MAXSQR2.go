package main

import "fmt"

type MATRIX [50][50]int

type node struct {
	row int
	column int
	next *node
}

func maxSquare(x MATRIX, n, key, row, column, width int) {
	var (
		count MATRIX
		list_head, last_node, p node
	)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if x[i][j] == key {
				count[i][j] = 1
				p = node{row:i, column:j}
				if list_head.next == nil {
					list_head = p
					last_node = p
				} else {
					last_node.next = &p
				}
			} else {
				count[i][j] = 0
			}
		}
	}

	for list_head.next != nil {
		last_node.next = nil
		p = list_head
		for p.next != nil {
			if count[p.row][p.column] == count[p.row+1][p.column] &&
			count[p.row][p.column] == count[p.row][p.column+1] &&
			count[p.row][p.column] == count[p.row+1][p.column+1] {
				count[p.row][p.column]++
				last_node = p
			} else {
				if last_node.next == nil {
					list_head = *(p.next)
				} else {
				last_node.next = p.next
				}
			}
			p = *(p.next)
		}
	}

	width = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if width < count[i][j] {
				width = count[i][j]
				row = i
				column = j
			}
		}
	}
}

func main() {
	var x = MATRIX{
		{ 1, 1, 1, 0, 0, 0, 0, 0},
		{ 0, 1, 1, 1, 1, 0, 0, 0},
		{ 0, 0, 1, 1, 1, 1, 1, 0},
		{ 0, 0, 0, 1, 1, 1, 1, 1},
		{ 0, 1, 1, 1, 1, 1, 1, 1},
		{ 0, 1, 1, 1, 1, 0, 1, 1},
		{ 0, 0, 1, 1, 1, 0, 0, 0},
		{ 0, 0, 1, 1, 1, 1, 0, 0},
	}

	var n int = 8
	var (
		row int = 1
		column int
		width int
		key int
	)

	fmt.Println("Find Maximum Size Square of Same Data")
    fmt.Println("=====================================")
	fmt.Println("Given Matrix : ")
	for row = 0; row < n; row++ {
		for column = 0; column < n; column++ {
			fmt.Printf("%d ", x[row][column])
		}
		fmt.Println()
	}

	fmt.Printf("Key to be searched = %d\n", key);

    maxSquare(x, n, key, row, column, width);

    fmt.Printf("Maximum Square is located at x[%d][%d] with size %d\n", row, column, width)
}