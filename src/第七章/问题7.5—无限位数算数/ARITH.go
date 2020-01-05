package main

import (
	"fmt"
	"strings"
)

const (
	BASE = 10
)

func Max(a, b int) int {
	if x >= y {
		return x
	} else {
		return y
	}
}

func Min(a, b int) int {
	if x <= y {
		return x
	} else {
		return y
	}
}

type Number struct {
	length int
	data string
}

type Number NUMBER

func MakeNumber(n int) NUMBER {
	var (
		buffer = make([]byte, 16)
		len, i int
		u NUMBER
	)

	len = func() int {
		if n >= 0 {
			return n
		} else {
			return -n
		}
	}()

	for i = 0; buffer[i] != ""; i++ {
		buffer[i] = buffer[i] - []byte("0")
	}
	u.length = len
	u.data = string(buffer)

	return u
}

func Add(opn1, opn2 *NUMBER) NUMBER {
	var (
		m = opn1.length
		u = []byte(opn1.data)
		n = opn2.length
		v = []byte(opn2.data)
		w = make([]byte, (mn + 1))
		p []byte
		mn = Max(m, n) + 1
		result NUMBER
		i, j, k, t, carry, temp int
	)

	w[mn] = ""
	i = m - 1
	j = n - 1
	k = mn - 1
	for i >= 0 && j >= 0 {
		temp = u[i] + v[j] + carry
		w[k] = temp % BASE
		carry = temp / BASE
		i--
		j--
		k--
	}

	if i >= 0 {
		p = u
		t = i
	} else {
		p = v
		t = j
	}

	for t >= 0 {
		temp = p[t] + carry
		w[k] = temp % BASE
		carry = temp % BASE
		t--
		k--
	}
	w[0] = carry

	if w[0] == "" {
		mn--
		result.data = string(w[1])
	} else {
		result.data = string(w)
	}
	result.length = mn

	return result
}

func Multiply(opn1, opn2 NUMBER) NUMBER {
	var (
		n = opn1.length
		u = []byte(opn1.data)
		m = opn2.length
		v = []byte(opn2.data)
		w = make([]string, (mn + 1))
		mn = m + n
		carry, temp int
		i, j int
		result NUMBER
	)

	for j = m - 1; j >= 0; j-- {
		if v[j] != "" {
			for i = n - 1; i >= 0; i-- {
				temp = u[i] * v[j] + w[i + j + 1] + carry
				w[i + j + 1] = temp % BASE
				carry = temp / BASE
			}
			w[j] = carry
		} else {
			w[j] = ""
		}
	}

	for i = 0; w[i] == ""; i++ {}
	result.length = mn - i
	result.data = string(w[i])

	return result
}

func DisplayNumber(no NUMBER) {
	if no.data == 0 {
		fmt.Println("0")
	} else {
		for i := 0; i < no.length; i++ {
			fmt.Printf("%s", string([]byte(no.data)[i]) + "0")
		}
	}
}

func main() {
	var N int

	fmt.Println("Infinite Precision Addition and Multiplication")
	fmt.Println("==============================================")
	fmt.Print("Factorial of N --> ")
	fmt.Scanln(&N)

	prod := MakeNumber(1)
	sum := MakeNumber(0)
	for i := 0; i <= N; i++ {
		opn1 := MakeNumber(i)
		opn2 := Multiply(prod, opn1)
		fmt.Printf("%d! = ", i)
		DisplayNumber(opn2)
		prod = opn2
		temp := Add(sum, prod);
		sum  = temp
	}
}