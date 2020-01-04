package main

import (
	"fmt"
	"unicode"
	"math"
)

const (
	OVERFLOW = 1
	UNDERFLOW = -1
	NO_ERROR = 0
)

func XAtoi(s []rune, result *int) int {
	*result = 0

	for i := 0; i < len(s) && unicode.IsDigit(s[i]); i++ {
		digitValue := int(s[i] - []rune("0")[0])
		if *result <= (math.MaxInt64 - digitValue) / 10 {
			*result = 10 * *result + digitValue
		} else {
			return OVERFLOW
		}
	}

	return NO_ERROR
}

func main() {
	var (
		result int
		s1 = "12345"
		s2 = "45678"
		s3 = "32768"
	)

	fmt.Println("atoi() Function with Overflow Detection")
	fmt.Println("=======================================")
	fmt.Printf("String #1 = \"%s\"\n", s1)
	if XAtoi([]rune(s1), &result) == NO_ERROR {
		fmt.Printf("Result    = %d\n", result)
	} else {
		fmt.Println("Result    = OVERFLOW")
	}
	fmt.Printf("String #2 = \"%s\"\n", s2)
	if XAtoi([]rune(s2), &result) == NO_ERROR {
		fmt.Printf("Result    = %d\n", result)
	} else {
		fmt.Println("Result    = OVERFLOW")
	}
	fmt.Printf("String #3 = \"%s\"\n", s3)
	if XAtoi([]rune(s3), &result) == NO_ERROR {
		fmt.Printf("Result    = %d\n", result)
	} else {
		fmt.Println("Result    = OVERFLOW")
	}
}