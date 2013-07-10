package main

import (
	"fmt";
)

// Minimum of a slice of int arguments
func MinIntS(v []int) (m int) {
	if len(v) > 0 {
		m = v[0]
	}
	for i := 1; i < len(v); i++ {
		if v[i] < m {
			m = v[i]
		}
	}
	return
}

// Minimum of a variable number of int arguments
func MinIntV(v1 int, vn ...int) (m int) {
	m = v1
	for i := 0; i < len(vn); i++ {
		if vn[i] < m {
			m = vn[i]
		}
	}
	return
}

// dp: word dist
func calc_dist(a, b string) int{
	var la, lb int = len(a), len(b)
	// define a marix
	var c = make([][]int, la + 1);
	for i := 0; i <= la; i++ {
		c[i] = make([]int, lb + 1)
	}

	for i := 0; i < la; i++ {
		c[i][lb] = la - i
	}
	for j := 0; j < lb; j++ {
		c[la][j] = lb - j
	}

	c[la][lb] = 0

	for i := la - 1; i >= 0; i-- {
		for j := lb - 1; j >= 0; j-- {
			if (b[j] == a[i]) {
				c[i][j] = c[i + 1][j + 1];
			} else {
				c[i][j] = MinIntV(c[i][j + 1], c[i + 1][j], c[i + 1][j + 1]) + 1
			}
		}
	}

	for i := 0; i <= la; i++ {
		fmt.Println(c[i])
	}

	return c[0][0]
}

func main() {
	var a, b string = "hello", "world"

	ret := calc_dist(a, b);
	fmt.Println("Result", a, b, ret)
}

