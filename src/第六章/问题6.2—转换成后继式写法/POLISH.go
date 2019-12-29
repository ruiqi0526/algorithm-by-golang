package main

import (
	"fmt"
	//"strings"
	"unicode"
	"os"
)

const (
	BOTTOM = iota
	EOL 
	LEFT_PAR 
	RIGHT_PAR 
	PLUS_MINUS 
	MUL_DIV 
)

var (
	stack = make([]string, 100)
	code = make([]string, 100)
	top int
)

func Initial() {
	top = 0
	code[top] = string(BOTTOM)
}

func StackTop() string {
	return code[top]
}

func Push(operator, opr_code string) {
	top++
	if top == 100 {
		fmt.Println("*** ERROR *** Stack Overflow")
		os.Exit(1)
	} else {
		stack[top] = operator
		code[top] = opr_code
	}
}

func Pop() string {
	if top == 0 {
		fmt.Println("*** ERROR *** Stack Underflow")
		os.Exit(1)
	} else {
		top--
	}
	return stack[top]
}

func main() {
	var p string
	var opr rune

	fmt.Println("Compile to Reversed Polish Program")
	fmt.Println("==================================")
	fmt.Println("Your Input --> ")
	fmt.Scanln(&p)
	fmt.Println("Answer is ---> ")

	Initial()
	for _, v := range []rune(p) {
		if unicode.IsLetter(v) {
			fmt.Printf("%v ", v)
		} else if string(v) == "(" {
			Push(string(v), string(LEFT_PAR))
		} else if unicode.IsSpace(v) {
			switch string(v) {
				case "+":
				case "-": 
					opr = PLUS_MINUS
				case "*":
				case "/":
					opr = MUL_DIV
				case ")":
					opr = RIGHT_PAR
				case "":
					opr = EOL
				default:
					fmt.Println("*** Unrecognizable char ***")
					os.Exit(0)
			}
			t := []rune(StackTop())[0]
			for t >= opr {
				fmt.Printf("%s ", Pop())
			}
			if t == LEFT_PAR && opr == RIGHT_PAR {
				Pop()
			} else if opr == EOL {
				os.Exit(1)
			} else {
				Push(string(v), string(opr))
			}
		}
	}
}