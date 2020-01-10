package main

import (
	"fmt"
)

type Node struct {
	vertex int
	next *Node
}

var (
	connect [30][30]int
	deg [30]int
	n int
	trail Node
)

func main() {
	fmt.Println("Euler Trail Program")
	fmt.Println("===================")
	
	readIn()
	trail = Euler()
	Display()
}

func readIn() {
	var (
		i, j int
	)

	fmt.Scanln(&n)
	for i = 0; i < n; i++ {
		deg[i] = 0
		for j = 0; j < n; j++ {
			connect[i][j] = 0
		}
	}
	fmt.Scanln(&i, &j)
	for i != 0 && j != 0 {
		if i != j {
			connect[i + 1][j + 1]++
			connect[j + 1][i + 1]++
			deg[i + 1]++
			deg[j + 1]++
		} else {
			fmt.Println("*** ERROR *** A loop found.  Data ignored")
		}
		fmt.Scanln(&i, &j)
	}
}

func Euler() Node {
	var (
		current Node
		head, tail Node
		p1, p2 Node
		VTX int
	)

	current = tail
	for {
		VTX = FindNext(current)()
		if VTX != -1 {
			FindTrail(VTX, p1, p2)
			p2.next = current.next
			current.next = &p1
			current = p1
		} else {
			break
		}
	}

	return head
}

func Prepare(first, last Node) int {
	var (
		p1, p2 Node
		no, oddNo, i int
		odd [2]int
	)

	for i = 0; i < n; i++ {
		if deg[i] % 2 != 0 {
			oddNo ++
			if no < 2 {
				odd[no] = i
				no++
			}
		}
	}
	if oddNo > 2 {
		fmt.Println("*** ERROR *** too many odd degree vertices.")
		return -1
	}
	if oddNo == 2 {
		connect[odd[0]][odd[1]]++
		connect[odd[1]][odd[0]]++
		deg[odd[0]]--
		deg[odd[1]]--
		p1.vertex = odd[0]
		p1.next = &p2
		p2.vertex = odd[1]
		first = p1
		last = p2
	} else {
		p1.vertex = 0
		last = p1
		first = p1
	}

	return 1
}

func FindNext(p Node) func() int {
	for p != (Node{}) && deg[p.vertex] == 0 {
		p = *(p.next)
	}

	return func() int {
		if p == (Node{}) {
			return -1
		} else {
			return p.vertex
		}
	}
}

func FindTrail(start int, head Node, tail Node) {
	var first, last, ptr Node
	var i int

	p := start
	for {
		for i = 0; i < n && connect[p][i] == 0; i++ {}
		if i < n {
			connect[p][i]--
			connect[i][p]--
			deg[p]--
			deg[i]--
			ptr.vertex = i
			if first == (Node{}) {
				last = ptr
				first = last
			} else {
				last.next = &ptr
				last = ptr
			}
			p = i
		} else {
			break
		}
	}
	head = first
	tail = last
}

func Display() {
	ptr := trail
	i := 0

	if trail == (Node{}) {
		return
	} 
	fmt.Println("An Euler Trail has been Found :")
	for *(ptr.next) != (Node{}) {
		if i % 15 == 0 {
			fmt.Println()
		}
		fmt.Printf("%2d->", ptr.vertex + 1)
		ptr = *(ptr.next)
		i++
	}
	if i % 15 == 0 {
		fmt.Println()
	}
	fmt.Printf("%2d", ptr.vertex + 1)
}