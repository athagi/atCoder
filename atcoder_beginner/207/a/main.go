package main

import (
	"fmt"
)

func main() {
	var A, B, C int
	fmt.Scan(&A, &B, &C)

	res := 0
	if A+B > res {
		res = A + B
	}

	if B+C > res {
		res = B + C
	}

	if C+A > res {
		res = C + A
	}

	fmt.Println(res)

}
