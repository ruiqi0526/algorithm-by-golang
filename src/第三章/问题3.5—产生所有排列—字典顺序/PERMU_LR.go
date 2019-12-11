package main

import "fmt"

var perm [20]int

func SWAP(a, b int) (int, int) {
	return b, a
}

func REVERSE(a, b int){
	j := b
	for i := a; i < j; i++ {
		j--
		perm[j], perm[i] = SWAP(perm[i], perm[j])
	}
}

func DISPLAY(n int) {
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", perm[i])
	}
}

func again(perm [20]int, L, R int) {
	var i = R

	for {
		if R - L + 1 > 2 {
			again(perm, L + 1, R)
			DISPLAY(R)
		}

		if i > L {
			perm[i], perm[L] = SWAP(perm[L], perm[i])
			REVERSE(L + 1, R)
			DISPLAY(R)
			i--
		}else{
			break
		}
	}
}

func permut(perm [20]int, n int){
	for i := 0; i < n; i++ {
		perm[i] = i + 1
	}
	again(perm, 0, n - 1)
}

func main() {
	var n int
	fmt.Scanln(&n)
	permut(perm, n)
}