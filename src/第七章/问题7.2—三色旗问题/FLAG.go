package main

import (
	"fmt"
	"strings"
)

const (
	BLUE = "b"
	WHITE = "w"
	RED = "r"
)

func SWAP(color []string, x, y int) {
	color[x], color[y] = color[y], color[x]
}

func DutchFlag(color []string) {
	var (
		white, blue int
		red = len(color) - 1
	)

	for white <= red {
		if color[white] == WHITE {
			white++
		} else if color[white] == BLUE {
			SWAP(color, blue, white)
			blue++
			white++
		} else {
			for white < red && color[red] == RED {
				red--
			}
			SWAP(color, red, white)
			red--
		}
	}
}

func main() {
	var flag string

	fmt.Println("Dutch National Flag Problem")
	fmt.Println("===========================")
	fmt.Println("Input a String of Color Tokens (b, w and r)")
	fmt.Scanln(&flag)

	flagArray := strings.Split(flag, "")
	DutchFlag(flagArray)

	fmt.Println("Rearranged Flag is :")
	fmt.Printf("%s\n", strings.Join(flagArray, ""))
}