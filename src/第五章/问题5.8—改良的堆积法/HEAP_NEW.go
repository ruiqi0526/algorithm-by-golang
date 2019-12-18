package heap_new

const (
	YES = 1
	NO = 0
)

func Sort(old_x []int, n int) {
	var x = make([]int, len(old_x))
	var temp int

	for i := n / 2; i >= 1; i-- {
		fix_heaq(x, i, x[i], n)
	}

	for size := n - 1; size >= 1; size-- {
		temp = x[size + 1]
		x[size + 1] = x[1]
		fix_heaq(x, 1, temp, size)
	}
}

func fix_heaq(x []int, root, key, bound int) {
	son := root + root
	level := 1
	for son < bound {
		son = func() int {
			if x[son] < x[son + 1] {
				return (son + 1) << 1
			} else {
				return son << 1
			}
		 }()
		 level++
	}

	if son > bound {
		level--
		son >>= 1
	}

	top := level--
	bottom := 0
	for top > bottom {
		mid := (top + bottom) / 2
		if mid <= x[son >> mid] {
			top = mid
		} else {
			bottom = mid + 1
		}
	}

	for mid := level - 1; mid >= top; mid-- {
		x[son >> (mid + 1)] = x[son >> mid]
	}
	x[son >> top] = key
}