package qsort_l

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

func Sort(x []int, n int) {
	var (
		first node
		last node
		p, q node
	)

	for i := 0; i < n; i++ {
		p.data = x[i]
		insert(first, last, p)
	}
	list_quicksort(first, last)

	p = first
	for i:= 0; i < n; i++ {
		x[i] = p.data
		q = p
		p = p.next
	}
}

func list_quicksort(first, last node) {
	var (
		less_first node
		less_last node
		equal_first node
		equal_last node
		greater_first node
		greater_last node
		p node
	)

	pivot := first.data
	p = insert(equal_first, equal_last, first)
	for p != (node{}) {
		if p.data < pivot {
			p = insert(less_first, less_last, p)
		} else if p.data > pivot {
			p = insert(greater_first, greater_last, p)
		} else {
			p = insert(equal_first, equal_last, p)
		}
	}

	join(less_first, less_last, equal_first, equal_last)
	join(less_first, less_last, greater_first, greater_last)

	first = less_first
	last = less_last
}

func insert(first, last, work node) node {
	var p node

	if first == (node{}) {
		last = work
		first = last
	} else {
		last.next = work
		last = work
	}

	p = work.next
	work.next = (node{})
	return p
}

func join(first, last, head, tail node) {
	if first == (node{}) {
		first = head
		last = tail
	} else if head != (node{}) {
		last.next = head
		last = tail
	}
}