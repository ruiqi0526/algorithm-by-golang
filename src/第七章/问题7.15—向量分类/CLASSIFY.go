package main

import (
	"fmt"
)

type Node struct {
	row int
	next *Node
}

var classNo int
var classList []Node

func Classify(data [][]int, m, n int, class []int) {
	var classList = make([]Node, m)
	
	classNo = 0
	var nodeList Node
	var tail Node
	var p Node

	nodeList.row = 0
	for i := 1; i < m; i++ {
		p.row = i
		tail.next = &p
		tail = *(tail.next)
	} 

	Group(nodeList, 0, n, data)

	for i := 0; i < classNo; i++ {
		for p = classList[i]; p != Node{}; p = *(p.next) {
			class[p.row] = i
		}
	}

	for nodeList != Node{} {
		p = nodeList
		nodeList = *(nodeList.next)
	}
}

func Group(list Node, column, n int, data [][]int) {
	var other Node

	if column >= n {
		classList[classNo] = list
		classNo++
	} else {
		other = Dispatch(list, column, data)
		Group(list, column + 1, n, data)
		if other != Node{} {
			Group(other, column, n, data)
		}
	}
}

func Dispatch(head Node, column int, data [][]int) Node {
	var (
		first Node
		last Node
		p Node
	)
	
	value := data[head.row][column]
	p = *(head.next)

	for ; p != Node{}; p = *(head.next) {
		if data[p.row][column] != value {
			head.next = p.next
			*(p.next) = Node{}
			if first == Node{} {
				last = p
				first = last
			} else {
				last.next = &p
				last = *(last.next)
			}
		} else {
			head = *(head.next)
		}
 	}

	return first
}

func main() {
	var data = [][]int{
		{ 3, 5, 7, 9},  { 4, 3, 7, 5},  { 3, 5, 7, 8},
		{ 2, 1, 4, 6},  { 3, 5, 7, 9},  { 1, 3, 5, 9},
		{ 4, 3, 7, 5},  { 2, 4, 6, 8},  { 2, 1, 4, 6},
		{ 4, 3, 1, 7},  { 3, 5, 7, 9},  { 1, 2, 3, 4},
		{ 1, 2, 3, 4},  { 9, 7, 5, 3},  { 2, 4, 6, 8},
	}
	var m = 15
	var n = 4
	var class = make([]int, 20)

	fmt.Println("Classification of Vectors")
	fmt.Println("=========================")

	Classify(data, m, n, class)
	fmt.Println("Classified Vectors with Code :")
	fmt.Println(" Code   Input Vectors")
	fmt.Println("-----   -------------")
	for i := 0; i < m; i++ {
		fmt.Printf("%5d - ", class[i])
		for j := 0; j < n; j++ {
			fmt.Printf("%3d\n", data[i][j])
		}
	} 
}