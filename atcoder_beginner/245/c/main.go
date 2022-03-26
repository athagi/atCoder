package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	A := make([]int, N)
	B := make([]int, N)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i := 0; i < N; i++ {
		sc.Scan()
		tmp, _ := strconv.Atoi(sc.Text())
		A[i] = tmp
	}

	for i := 0; i < N; i++ {
		sc.Scan()
		tmp, _ := strconv.Atoi(sc.Text())
		B[i] = tmp
	}

	a := A[N-1]
	b := B[N-1]
	for i := N - 2; i >= 0; i-- {
		x := A[i]
		y := B[i]
		var tmpx, tmpy int

		if a > 0 {
			if abs(a, x) <= K {
				tmpx = x
			} else {
				tmpx = -1
			}
			if abs(a, y) <= K {
				tmpy = y
			} else {
				tmpy = -1
			}
		}
		if b > 0 {
			if abs(b, x) <= K {
				tmpx = x
			}
			if abs(b, y) <= K {
				tmpy = y
			}
		}

		a = tmpx
		b = tmpy
	}
	res := "No"
	if a > 0 || b > 0 {
		res = "Yes"
	}
	fmt.Println(res)
}

func abs(a, b int) int {
	if a < 0 || b < 0 {
		return -1
	}
	res := a - b
	if res < 0 {
		res *= -1
	}
	return res
}
