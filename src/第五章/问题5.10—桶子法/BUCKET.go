package bucket

const (
	LENGTH = 4
	BUCKET_NO = 16
	LOOP_COUNT = 16
)

type node struct {
	data int
	x int
	next node
}

type cell struct {
	first node
	last node
}

var bucker = [BUCKET_NO]cell
var mask int = 0x0f
var head cell

func Sort(input []int, n int) {

	head = listgen(input, n)

	for count := 1; count <= LOOP_COUNT; count++ {
		for i := 0; i < BUCKET_NO; i++ {
			bucket[i].last = nil
			bucker[i].first = bucket[i].last
		}
		distribute()
		recollect()
	}
	putback(input)
}

func listgen(input []int, n int) node {
	var (
		first node
		last node
		temp node
	)

	first.x = input[0]
	first.data = input[0]
	first.next = nil
	for i := 1; i < n; i++ {
		temp.x = input[i]
		temp.data = input[i]
		temp.next = nil
		last.next = temp
		last = temp
	}
	return first
}

func distribute() {
	var (
		temp node
		index int
	)

	for head != nil {
		index = head.x & mask
		head.x >>= LENGTH
		if bucket[index].first == nil {
			bucket[index].first = head
			bucket[index].last = head
		} else {
			bucket[index].last.next = head
			bucket[index].last = head
		}
		temp = head.next
		if head.next != nil {
			head.next.next = nil
		}
		head = temp
	}
}

func recollect() {
	i := 0
	if i < BUCKET_NO && bucket[i].first == nil {
		i++
	}

	head = bucket[i].first
	for j := i + 1; j < BUCKET_NO; j++ {
		if bucket[j].first != nil {
			bucket[i].last.next = bucket[j].first
			i = j
		}
	}
}

func putback(input []int) {
	var temp node

	for i := 0; head != nil; i++ {
		input[i] = head.data
		temp = head.next
		head = temp
	}
}