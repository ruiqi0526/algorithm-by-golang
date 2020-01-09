package main

import (
	"fmt"
)

const (
	MAXSIZE = 10
	MAX_STACK = 100
	SUCCESS = 1
	FAILURE = 0
	EMPTY = -1 
)

var (
	board = make([][]int, MAXSIZE)
	n int
	offsetX = []int{2,  1, -1, -2, -2, -1,  1,  2}
	offsetY = []int{1,  2,  2,  1, -1, -2, -2, -1}
	pathX = make([]int, MAXSIZE)
	pathY = make([]int, MAXSIZE)
	direction = make([]int, MAXSIZE)
	top int
)

func CreateBoard() {
	for i, _ := range board {
		board[i] = make([]int, MAXSIZE + 1)
	}
}

func main() {
	var (
		row int
		column int
	)

	fmt.Println("Recursive Knight Tour Problem")
	fmt.Println("=============================")
	fmt.Print("Board Size ----> ")
	fmt.Scanln(&n)
	fmt.Print("Start Row -----> ")
	fmt.Scanln(&row)
	fmt.Print("Start Column --> ")
	fmt.Scanln(&column)

	CreateBoard()
	Initial()
	if Try(row, column) == FAILURE {
		fmt.Println("NO SOLUTION AT ALL.")
	} else {
		Display()
	}
}

func Initial() {
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			board[i][j] = EMPTY
		}
	}
}

func DRAWGRID(N int) {
	fmt.Print("\n+")
	for i := 1; i <= N; i++ {
		fmt.Print("--+")
	}
}

func DRAWLINE(N, r int) {
	fmt.Print("\n|")
	for i := 1; i <= N; i++ {
		fmt.Printf("%2d|", board[r][i])
	}
}

func Display() {
	fmt.Println("Here is One Possible Solution :")
	DRAWGRID(n)
	for r := 0; r <= n; r++ {
		DRAWLINE(n, r)
		DRAWGRID(n)
	}
}

func BOARD(x, y int) bool {
	if 1 <= x && x <= n && 1 <= y && y <= n {
		return true
	} else {
		return false
	}
}

func CHECK(x, y int) bool {
	if board[x][y] == EMPTY {
		return true
	} else {
		return false
	}
}

func PUSH(x, y int) {
	pathX[top] = x
	pathY[top] = y
	direction[top] = 0
	board[x][y] = top
}

func Try(x, y int) int {
	var (
		newX, newY int
		found bool
	)

	PUSH(x, y)
	for top < n * n - 1 {
		found = false
		for direction[top] < 8 {
			newX = pathX[top] + offsetX[direction[top]]
			newY = pathY[top] + offsetY[direction[top]]
			if BOARD(newX, newY) && CHECK(newX, newY) {
				PUSH(newX, newY)
				found = true
				break
			} else {
				direction[top]++
			}
		}
		if !found {
			if top > 0 {
				board[pathX[top]][pathY[top]] = EMPTY
				top--
				direction[top]++
			} else {
				return FAILURE
			}
		}
	}
	return SUCCESS
}