package main

import (
	"fmt"
	"strings"
)

func LongestCommonSubSequence(a, b, result []string) {
	var (
		m = len(a)
		n = len(b)
		d = make([][]int, 8 * (m + 1))
		count, i, j int
	)

	for i = 0; i < 8 * (m + 1); i++ {
		d[i] = make([]int, 8 * (m + 1) * (n + 1))
	}

	for i = 1; i <= m; i++ {
		d[i] = d[i + n]
	}

	d[0][0] = 0
	for i = 1; i <= m; i++ {
		d[i][0] = 0
	}
	for j = 1; j <= n; j++ {
		d[0][j] = 0
	}

	for i = 1; i <= m; i++ {
		for j = 1; j <= n; j++ {
			if a[i - 1] == b[j - 1] {
				d[i][j] = d[i - 1][j - 1] + 1
			} else if d[i][j - 1] > d[i - 1][j] {
				d[i][j] = d[i][j - 1]
			} else {
				d[i][j] = d[i - 1][j]
			}
		}
	}

	count = d[m][n]
	i = m
	j = n
	for i != 0 && j != 0 {
		if d[i][j] == d[i - 1][j] {
			i--
		} else if d[i][j] == d[i][j - 1] {
			j--
		} else {
			count--
			result[count] = a[i - 1]
			i--
			j--
		}
	}
}

func main() {
	var (
		a = "ABXZWTP01=+C*US"
		b = "TX013W+-P12=+CCTXU"
		c = make([]string, 50)
	)

	fmt.Println("Longest Common Subsequence Program")
	fmt.Println("==================================")
	fmt.Printf("First  String : %s\n", a)
	fmt.Printf("Second String : %s\n", b)
	aArray := strings.Split(a, "")
	bArray := strings.Split(b, "")
	LongestCommonSubSequence(aArray, bArray, c)
	fmt.Printf("L.C.S. String : %s\n", strings.Join(c, ""))
}