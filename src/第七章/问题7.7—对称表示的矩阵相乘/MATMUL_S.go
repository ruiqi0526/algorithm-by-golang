package main

import (
	"fmt"
)

func StwoDToOneD(TwoD [][]int, n int, OneD []int) {
	var index int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			OneD[index] = TwoD[i][j]
			index++
		}
	}
}

func OneDToStwoD(OneD []int, TwoD [][]int, n int) {
	index := n * (n + 1) / 2 - 1

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= i; j-- {
			TwoD[j][i] = OneD[index]
			TwoD[i][j] = TwoD[j][i]
			index--
		}
	}
}

func GAP(k, i, gap int) int {
	if k <= i {
		return gap
	} else {
		return 1
	}
}

func Smatmul(a, b, c []int, n int) {
	var (
		index int
		i, j, k int
		p, q, gap int
		sum int
	)

	for i = 0; i < n; i++ {
		for j = 0; j < n; j++ {
			gap = n - 1
			p = i
			q = j
			for k = 0; k < n; k++ {
				sum += a[p] * b[q]
				q += GAP(k, i, gap)
				q += GAP(k, j, gap)
				gap--
			}
			c[index] = sum
			index++
		}
	}
}

func OneDToTwoD(OneD []int, TwoD [][]int, m, n int) {
	index := m * n - 1

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			TwoD[i][j] = OneD[index]
			index--
		}
	}
}

func main() {
	var A = [][]int{
		{ 3, 5, 7, 9, 2},
		{ 5, 5, 1, 7, 3},
		{ 7, 1, 1, 4, 5},
		{ 9, 7, 4, 3, 1},
		{ 2, 3, 5, 1, 4},
	}
	var B = [][]int{
		{ 1, 3, 5, 7, 9},
		{ 3, 4, 6, 8, 1},
		{ 5, 6, 9, 2, 4},
		{ 7, 8, 2, 1, 2},
		{ 9, 1, 4, 2, 5},
	}
	var C = make([][]int, 20)
	for i, _ := range C {
		C[i] = make([]int, 20)
	}

	var N = 5

	fmt.Println("One Dimensional Symmetric Matrix Multiplication")
	fmt.Println("===============================================")
	fmt.Println("Matrix A[][] :")
	fmt.Printf("%v\n", A)

	AA := A[0]
	for _, v := range A[1:] {
		AA = append(AA, v...)
	}
	StwoDToOneD(A, N, AA)
	fmt.Println("Transformed One Dimensional Array A[] :")
	fmt.Printf("%v\n", AA)

	fmt.Println("Matrix B[][]:")
	fmt.Printf("%v\n", B)

	BB := B[0]
	for _, v := range B[1:] {
		BB = append(BB, v...)
	}
	StwoDToOneD(B, N, BB)
	fmt.Println("Transformed One Dimensional Array B[] :")
	fmt.Printf("%v\n", BB)

	CC := C[0]
	for _, v := range C[1:] {
		CC = append(CC, v...)
	}
	Smatmul(AA, BB, CC, N)
	fmt.Println("Linear Form of A[][] * B[][] :")
	fmt.Printf("%v\n", CC)

	OneDToTwoD(CC, C, N, N)
	fmt.Println("Matrix C[][]:")
	fmt.Printf("%v\n", C)
}