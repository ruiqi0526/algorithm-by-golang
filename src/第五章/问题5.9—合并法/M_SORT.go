package m_sort

const SIZE = 8

func merge(x []int, work []int, first, mid, last int) {
	var (
		in1 = first
		in2 = mid + 1
		out = first
	)

	for in1 <= mid && in2 <= last {
		out++
		work[out] = func() int {
			if x[in1] <= x[in2] {
				in1++
				return x[in1]
			} else {
				in2
				return x[in2]
			}
		}()
	}
	for in1 <= mid {
		out++
		in1++
		work[out] = x[in1]
	}
	for in2 <= last {
		out++
		in2++
		work[out] = x[in2]
	}
}

func mergesort(x []int, work []int, first, last int) {
	if first < last {
		mid := (first + last) / 2
		mergesort(x, work, first, mid)
		mergesort(x, work, mid + 1, last)
		merge(x, work, first, mid, last)
		x[first] = work[first][:(last - first + 1) * SIZE]
	}
}

func Sort(x []int, n int) {
	var work = make([]int, n * SIZE)
	mergesort(x, work, 0, n - 1)
}