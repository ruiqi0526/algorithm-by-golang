package heapsort

const (
	YES int = 1
	NO int = 0
)

func Sort(old_x []int, n int) {
	var x = make([]int, len(old_x))
	var temp int

	for i := n / 2; i >= 1; i-- {
		fix_heaq(x, i, x[i], n)
	}

	for size := n; size >= 2; size-- {
		temp = x[1]
		fix_heaq(x, 1, x[size], size - 1)
		x[size] = temp
	}
}

func fix_heaq(x []int, root, key, bound int) {
	father := root
	son := father + father
	done := NO

	for son <= bound && !done {
		if son < bound && x[son + 1] > x[son] {
			son++
		}
		if key < x[son] {
			x[father] = x[son]
			father = son
		} else {
			done = YES
		}
		son = father + father
	}
	x[father] = key
}