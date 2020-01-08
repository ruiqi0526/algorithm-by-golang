package main

import (
	"fmt"
)

const (
	MAXSIZE = 20
	DIAG_SIZE = 39
	OCCUPIED = 0
	EMPTY = 1
)

var (
	pos = make([]int, MAXSIZE + 1)
	inRow = make([]int, MAXSIZE + 1)
	diag = make([]int, MAXSIZE + 1)
	backDiag = make([]int, MAXSIZE + 1)
	n, count int
)

func main() {
	fmt.Println("All Possible Solutions of N Queens' Problem")
	fmt.Println("===========================================")
	fmt.Print("Board Size (N > 3) --> ")
	fmt.Scanln(&n)

	Initial()
	Try(1)
	fmt.Printf("There are %d solutions in total.\n", count)
}

func Initial() {
	for i := 1; i <= n; i++ {
		inRow[i] = EMPTY
	}
	for i := 1; i <= n * 2 - 1; i++ {
		backDiag[i] = EMPTY
		diag[i] = EMPTY
	}

	fmt.Println("Solutions are represented the row # in a column");
	fmt.Println("  #")
	for i := 1; i <= n; i++ {
		fmt.Printf("%3d", i)
	}
	fmt.Println("---")
	for i := 1; i <= n; i++ {
		fmt.Println(" --")
	}
}

func Display() {
	fmt.Printf("%3d", count)
	for i := 1; i <= n; i++ {
		fmt.Printf("%3d\n", pos[i])
	}
}

func SAVE(r, c int) bool {
	if inRow[r] == 1 && backDiag[r + c - 1] == 1 && diag[n - (c - r)] == 1 {
		return true
	} else {
		return false
	}
}

func SET(r, c int) {
	pos[c] = r
	inRow[r] = OCCUPIED
	backDiag[r + c - 1] = OCCUPIED
	diag[n - (c - r)] = OCCUPIED
}

func RESET(r, c int) {
	inRow[r] = EMPTY
	backDiag[r + c - 1] = EMPTY
	diag[n - (c - r)] = EMPTY
}

func Try(column int) {
	for row := 1; row <= n; row++ {
		if SAVE(row, column) {
			SET(row, column)
			if column < n {
				Try(column + 1)
			} else {
				count++
				Display()
			}
			RESET(row, column)
		}
	}
}
