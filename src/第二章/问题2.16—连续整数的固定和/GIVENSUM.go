package main

import "fmt"

func main() {
	var left, right, sum, count, GIVEN int

	fmt.Println("Consecutive sum to a fixed given number")
    fmt.Println("=======================================")
	fmt.Print("Your number (> 0) please ---> ")
	fmt.Scanln(&GIVEN)

	sum = 0
	for right = 1; sum < GIVEN; right++ {
		sum += right
	}

	for left = 1; left <= GIVEN/2; right-- {
		if sum > GIVEN {
			left++
			sum -= left
		} else {
			if sum == GIVEN {
				fmt.Printf("%d = sum from %d to %d\n", GIVEN, left, right)
				count++
			}
			right++
			sum += right - 1
		}
	}

	if count > 0 {
		fmt.Printf("There are %d solutions in total.\n", count)
	} else {
		fmt.Println("Sorry, there is NO solution at all.")
	}

}