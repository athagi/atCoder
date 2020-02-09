package main

import (
	"fmt"
)

func main() {
	var S, T, U string
	var A, B int
	fmt.Scan(&S, &T)
	fmt.Scan(&A, &B)
	fmt.Scan(&U)

	if S == U {
		A = A - 1
	}
	if T == U {
		B = B - 1
	}
	fmt.Println(A, B)
}
