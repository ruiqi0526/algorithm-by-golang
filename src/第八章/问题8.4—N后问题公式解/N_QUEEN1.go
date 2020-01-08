package main

import (
	"fmt"
)

const (
	MAXSIZE = 30
	MARK = "Q"
)

var n, i int
var board = make([][]string, MAXSIZE + 1)

func MakeBoard() {
	for i = 0; i < MAXSIZE + 1; i++ {
		board[i] = make([]string, MAXSIZE + 1)
	}
}

func main() {
	MakeBoard()
	var funct = []func(){Class1, Class2, Class4, Class3, Class1, Class2}

	fmt.Println("One Solution for N Queens' Problem")
	fmt.Println("==================================")
	fmt.Print("Board Size Please (N > 3) --> ")
	fmt.Scanln(&n)

	if n > 3 {
		Initial()
		funct[n % 6]()
		Display()
	} else {
		fmt.Println("Illegal Board Size.")
	}
}

func Initial() {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			board[i][j] = " "
		}
	}
}

func DRAWGRID(N int) {
	fmt.Println("+")
	for i := 1; i <= N; i++ {
		fmt.Println("-+")
	}
}

func DRAWLINE(N, r int) {
	fmt.Println("|")
	for i := 1; i <= N; i++ {
		fmt.Printf("%s|\n", board[r][i])
	}
}

func Display() {
	DRAWGRID(n)
	for r := 1; r <= n; r++ {
		DRAWLINE(n, r)
		DRAWGRID(n)
	}
}

func Class1() {
	for i := 1; i <= n / 2; i++ {
		board[2 * i - 1][n / 2 + i] = MARK
		board[2 * i][i] = board[2 * i - 1][n / 2 + i]
	}
}

func Class2() {
	for i := 1; i <= (n - 1) / 2; i++ {
		board[2 * i][i] = MARK
	}
	for i := 1; i <= (n + 2) / 2; i++ {
		board[2 * i - 1][(n - 1) / 2 + i] = MARK
	}
}

func Class3() {
	for i := 1; i <= (n - 3) / 2; i++ {
		board[2 * i + 3][(n - 1) / 2 + i] = MARK
		board[2 * i + 2][i] = board[2 * i + 3][(n - 1) / 2 + i]
	}
	board[3][n] = MARK
	board[1][n-1] = MARK
	board[2][(n-2)/2] = MARK
}

func Class4() {
	if n > 8 {
		for i := 1; i <= 3; i++ {
			board[2 * i - 1][n / 2 - 2 + i] = MARK
		}
		board[6][n - 1] = MARK
		board[2][n] = MARK
		board[4][1] = MARK
		for i := 1; i <= n / 2; i++ {
			board[2 * i + 6][n / 2 + 1 + i] = MARK
			board[2 * i + 5][i + 1] = MARK
		}
	} else {
		for i := 0; i <= 4; i++ {
			board[2 * i][i] = MARK
		}
		board[1][6] = MARK
		board[3][5] = MARK
		board[5][8] = MARK
		board[7][7] = MARK
	}
}