package main

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	LINE_SIZE = 100
	STACK_BOTTOM = 0
	OPERAND = 1
	OPERATOR = 2
)

func IsOpr(opr string) bool {
	return opr == "+" || opr == "-" || opr == "*" || opr == "/"
}

func Compute(opr string, opn1, opn2 float64) float64 {
	var result float64

	switch opr {
		case "+":
			result = opn1 + opn2
		case "-":
			result = opn1 - opn2
		case "*":
			result = opn1 * opn2
		case "/":
			result = opn1 / opn2
	}
	return result
}

type Union struct {
	value float64
	operator string
}

type Item struct {
	store *Union
	Type int
}

var (
	stack = make([]Item, 100)
	top int
)

func Initial() {
	top = 0
	stack[top].Type = STACK_BOTTOM
}

func PushOpn(data float64) {
	top++
	stack[top].Type = OPERAND
	stack[top].store.value = data
}

func PushOpr(opr string) {
	stack[top].Type = OPERATOR
	stack[top].store.operator = opr
	top++
}

func PopOpn() float64 {
	if top > 0 {
		top--
	} else {
		top = 0
	}
	
	return stack[top].store.value
}

func PopOpr() string {
	if top > 0 {
		top--
	} else {
		top = 0
	}
	
	return stack[top].store.operator
}

func StackTop() int {
	return stack[top].Type
}

func StringToRune(s string) rune {
	return []rune(s)[0]
}

func main() {
	var (
		opn1, opn2 float64
		opr string
		s string
	)

	fmt.Println("Prefix Form Evaluator")
	fmt.Println("=====================")
	fmt.Print("Input --> ")
	fmt.Scanln(&s)

	Initial()
	p := strings.Split(s, "")
	for _, v := range p {
		if IsOpr(v) {
			PushOpr(v)
		} else if unicode.IsDigit(StringToRune(v)) {
			opn2 = float64(StringToRune(v) - StringToRune("0"))
			for StackTop() == OPERAND {
				opn1 = PopOpn()
				opr = PopOpr()
				opn2 = Compute(opr, opn1, opn2)
			}
			PushOpn(opn2)
		}
	}
	fmt.Printf("Result = %d\n", PopOpn())
}