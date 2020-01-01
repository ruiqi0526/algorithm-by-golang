package main

import "fmt"

const (
	OK = 1
	NO = 0
)

func SeqGen(seq []int, n int) {
	seq[0] = 1
	seq[1] = 2
	no := 1

	for no >= 0 {
		if Check(seq, no) == OK {
			if no == n - 1 {
				OutPut(seq, n)
				Change(seq, no)
			} else {
				Extend(seq, no)
			}
		} else {
			Change(seq, no)
		}
	}
}

func Change(seq []int, no int) {
	for ; seq[no] == 3; no-- {}
	seq[no]++
}

func Extend(seq []int, no int) {
	no++
	seq[no] = 1
}

func Check(seq []int, no int) int {
	half := (no + 1) / 2
	var i int

	for length := 1; length <= half; length++ {
		for i < length && no - i >= 0 && no - length - i >= 0 {
			if seq[no - i] == seq[no - length - 1] {
				i++
			}
		}
		if i >= length {
			return NO
		}
	}
	return OK
}

func OutPut(seq []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", seq[i])
	}
}

func main() {
	var (
		seq = make([]int, 10)
		n int
	)

	fmt.Println("Non-Repeated Sequence of Three Chars. Generator")
	fmt.Println("===============================================")
	fmt.Print("Sequence Length Please --> ")
	fmt.Scanln(&n)
	fmt.Print("Generated Sequence is shown below :", SeqGen(seq, n))
}