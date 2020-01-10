package main

import (
	"fmt"
)

func main() {
	var (
		numberOfDisk int
		pin = make([]int, 30)
		dir = make([]int, 2)
		disk, next, index int
		numberOfMoves, counter, i int
	)

	fmt.Println("Iterative Towers of Hanoi Program")
	fmt.Println("=================================")
	fmt.Print("How many disks (<=31) ? ")
	fmt.Scanln(&numberOfDisk)

	numberOfMoves = (0x01 << uint64(numberOfDisk)) - 1
	counter = 0

	if (numberOfDisk & 0x01) == 1 {
		dir[0] = 0
		dir[1] = 1
	} else {
		dir[0] = 1
		dir[1] = 0
	}

	for i = 1; i <= numberOfDisk; i++ {
		pin[i] = 1
	}

	fmt.Println("   Step   Disk #    From      To")
	fmt.Println("   -----------------------------")
	
	for i = 1; i <= numberOfMoves; i++ {
		disk = WhichDisk(&counter)
		index = disk & 0x01
		next = (pin[disk] + dir[index]) % 3 + 1
		fmt.Printf("%6d%8d%10d%8d\n", i, disk, pin[disk], next)
		pin[disk] = next
	}
}

func WhichDisk(counter *int) int {
	a := *counter
	b := a + 1
	*counter = b
	i := 0

	for c := a ^ b; c != 0; c >>= 1 {
		i++
	}

	return i
}