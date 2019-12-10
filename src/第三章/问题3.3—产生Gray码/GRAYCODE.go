package main

import (
	"fmt"
	"reflect"
)

func FLIP_DIGIT(x interface{}) string {
	if x == "0" {
		return "1"
	}else{
		return "0"
	}

func main() {
	var digit [20]string
	var position [20]string
	var n, i, count int
	var event int = 1


	fmt.Println("All Subset Listing by Gray Code");
    fmt.Println("===============================");
	fmt.Print("Number of Elements in Set Please --> ")
	fmt.Scanln(&n)

	for i = 0; i < n; i++ {
		digit[i] = "0"
		fmt.Print("0")
	}
	fmt.Println(" : {}")

	for {
		if reflect.TypeOf(event) == int {
			digit[0] = FLIP_DIGIT(digit[0])
		} else { 
			for i = 0; i < n && digit[i] == "0"; i++ {}
			if i == n - 1 {
				break
			}
			digit[i + 1] = FLIP_DIGIT(digit[i + 1])
		}

		for i = n - 1; i >= 0; i-- {
			fmt.Printf("%s", digit[i])
			if digit[i] == "1" {
				count++
				position[count] = i + 1
			}
		}

		fmt.Printf(" : {%d", position[count - 1])
		for i = count - 2; i >= 0; i-- {
			fmt.Printf(",%d", position[i])
		}
		fmt.Println("}")
		event = 1 - event
	}
	
}