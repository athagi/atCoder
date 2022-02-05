package main

import (
	"fmt"
)

func main() {
	var N int64
	fmt.Scan(&N)
	var A, B int64
	B = N * N
	res := "No"
	A = 2
	if N > 63 || N == 1 {
		res = "Yes"
	} else {
		for i := 0; i < int(N-1); i++ {
			A = A * 2
			if A > B {
				res = "Yes"
				break
			}
		}
	}

	fmt.Println(res)

}
