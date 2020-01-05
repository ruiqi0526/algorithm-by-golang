package main

import (
	"fmt"
)

func TwoDToOneD(TwoD [][]int, m, n int, OneD []int) {
	var index int

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			OneD[index] = TwoD[i][j]
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

func Matmul(a, b, c []int, L, M, N int) {
	var (
		index int
		i, j, k int
		p, pp, qq int
		sum int
	)

	for ; i < L; i++ {
		for ; j < M; j++ {
			pp = p
			qq = j
			for ; k < M ; k++ {
				sum += a[pp] * b[qq]
				pp++
				qq += N
			}
			c[index] = sum
			index++
		}
		p += M
	}
}

func main() {
	var A = [][]int{
		{ 3, 5, 7, 9},
		{ 2, 5, 1, 7},
		{ 8, 2, 1, 4},
		{ 6, 8, 3, 3},
	}

	var B = [][]int{
		{ 1, 3, 5},
		{ 2, 4, 6},
		{ 3, 7, 9},
		{ 4, 5, 8},
	}

	var C = make([][]int, 20)
	for i, _ := range C {
		C[i] = make([]int, 20)
	}

	var (
		L = 4
		M = 4
		N = 3
	)

	fmt.Println("One Dimensional Matrix Multiplication Program")
	fmt.Println("=============================================")

	fmt.Println("Matrix A[][] :")
	fmt.Printf("%v\n", A)
	AA := A[0]
	for _, v := range A[1:] {
		AA = append(AA, v...)
	}
	TwoDToOneD(A, L, M, AA)
	fmt.Println("Transformed One Dimensional Array A[] :")
	fmt.Printf("%v\n", AA)

	fmt.Println("Matrix B[][]:")
	fmt.Printf("%v\n", B)
	BB := B[0]
	for _, v := range B[1:] {
		BB = append(BB, v...)
	}
	TwoDToOneD(B, M, N, BB)	
	fmt.Println("Transformed One Dimensional Array B[] :")
	fmt.Printf("%v\n", BB)

	CC := C[0]
	for _, v := range C[1:] {
		CC = append(CC, v...)
	}
	Matmul(AA, BB, CC, L, M, N)
	fmt.Println("Linear Form of A[][] * B[][] :")
	fmt.Printf("%v\n", CC)

	OneDToTwoD(CC, C, L, N)
	fmt.Println("Matrix C[][]:")
	fmt.Printf("%v\n", C)
}