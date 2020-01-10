package main

import (
	"fmt"
	"strings"
)

const (
	MAXSIZE = 50
	OCCUPIED = 1
	UNOCCUPIED = 0
	YES = 1
	NO = 0
)

var (
	cell = make([][]int, MAXSIZE)
	workCopy = make([][]int, MAXSIZE)
	row, column, generations int
)

func MakeSlices() {
	for i, _ := range cell {
		cell[i] = make([]int, MAXSIZE)
	}
	for i, _ := range workCopy {
		workCopy[i] = make([]int, MAXSIZE)
	}
}

func main() {
	MakeSlices()
	readIn()
	GameOfLife()
}

func readIn() {
	var (
		maxRow, maxCol int
		colGap, rowGap int
		i, j int
	)

	fmt.Scanln(&generations, &row, &column) 
	line := string(generations) + string(row) + string(column)
	lineArray := strings.Split(line, "")
	for maxRow = 0; line != "   "; maxRow++ {
		for i = 0; lineArray[i] != ""; i++ {
			if lineArray[i] != " " {
				cell[maxRow][i] = OCCUPIED
			}
		}
		maxCol = func() int {
			if maxCol < i {
				return i
			} else {
				return maxCol
			}
		}()
	}
	rowGap = (row - maxRow) / 2
	colGap = (column - maxCol) / 2
	for i = maxCol + rowGap - 1; i >= rowGap; i-- {
		for j = maxCol + colGap - 1; j >= colGap; j-- {
			cell[i][j] = cell[i - rowGap][j - colGap]
		}
		for ; i >= 0; i-- {
			for j = 0; j < column; j++ {
				cell[i][j] = UNOCCUPIED
			}
		}
	}
}

func DrawBoarder(n int) {
	fmt.Print("\n+")
	for i := 0; i < n; i++ {
		fmt.Print("-")
	}
	fmt.Print("+")
}

func Display(genNo int) {
	if genNo == 0 {
		fmt.Printf("\n\nInitial Generation :\n")
	} else {
		fmt.Printf("\n\nGeneration %d :\n", genNo)
	}
	DrawBoarder(column)
	for i := 0; i < row; i++ {
		fmt.Printf("\n|")
		for j := 0; j < column; j++ {
			fmt.Printf("%s", func() string {
				if cell[i][j] == OCCUPIED {
					return "*"
				} else {
					return " "
				}
			}())
		}
		fmt.Print("|")
	}
	DrawBoarder(column)
}

func GameOfLife() {
	var (
		stable bool
		iter int
		top, bottom, left, right int
		neighbors int
		cellCount int
		done int
	)

	Display(0)
	for iter = 1; iter <= generations && done == 0; iter++ {
		stable = true
		for i := 0; i < row; i++ {
			top = func() int {
				if i == 0 {
					return 0
				} else {
					return i - 1
				}
			}()
			bottom = func() int {
				if i == row-1 {
					return row - 1
				} else {
					return i + 1
				}
			}()
			for j := 0; j < column; j++ {
				left = func() int {
					if j == 0 {
						return 0
					} else {
						return j - 1
					}
				}()
				right = func() int {
					if j == column-1 {
						return column - 1
					} else {
						return j + 1
					}
				}()
				for p := top; p <= bottom; p++ {
					for q := left; q <= right; q++ {
						neighbors += workCopy[p][q]
					}
				}
				neighbors -= workCopy[i][j]

				if workCopy[i][j] == OCCUPIED {
					if neighbors == 2 || neighbors == 3 {
						cell[i][j] = OCCUPIED
						cellCount++
					} else {
						cell[i][j] = UNOCCUPIED
					}
				} else if neighbors == 3 {
					cell[i][j] = OCCUPIED
					cellCount++
				} else {
					cell[i][j] = UNOCCUPIED
				}
				stable = stable && (workCopy[i][j] == cell[i][j])
			}
			if cellCount == 0 {
				fmt.Println("All cells die out.")
				done = 1
			} else if stable {
				fmt.Println("System enters a stable state.")
				done = 1
			} else {
				Display(iter)
			}
		}
	}
}