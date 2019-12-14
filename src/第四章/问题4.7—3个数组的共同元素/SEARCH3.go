package main

import "fmt"

var (
	FOUND = 1
	NOT_FOUND = 2
)

func search(x, y, z []int, X, Y, Z, XX, YY, ZZ int) int {
	for XX < X && YY < Y && ZZ < Z {
		if x[XX] < y[YY] {
			XX++
		} else if y[YY] < z[ZZ] {
			YY++
		} else if z[XX] < x[XX] {
			ZZ++
		} else {
			return FOUND
		}
	}
	return NOT_FOUND
}

func main() {
	var (
		x = []int{1, 3, 5,  7,  9, 11, 13, 15, 17, 19}
		y = []int{2, 4, 9, 10, 12, 14, 16, 18, 20, 21}
		z = []int{1, 2, 3,  4,  5,  6,  7,  8,  9, 10}
		X = len(x) / 8
		Y = len(y) / 8
		Z = len(z) / 8
		XX, YY, ZZ int
	)

	fmt.Println("Search for a Common Element from Three Arrays")
	fmt.Println("=============================================")
	fmt.Println("First Array :")
	for XX = 0; XX < X; XX++ {
		fmt.Printf("%d", x[XX])
	}
	fmt.Println("\nSecond Array :")
	for YY = 0; YY < Y; YY++ {
		fmt.Printf("%d", y[YY])
	}
	fmt.Println("\nThird Array :")
	for ZZ = 0; ZZ < Z; ZZ++ {
		fmt.Printf("%d", z[ZZ])
	}

	if search(x, y, z, X, Y, Z, XX, YY, ZZ) > 0 {
		fmt.Printf("%d is common to x[%d], y[%d] and z[%d]\n", x[XX], XX, YY, ZZ)
	} else {
		fmt.Println("NO COMMON ELEMENT FOUND.")
	}
}