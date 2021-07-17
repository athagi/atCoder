package main

import (
	"fmt"
)

func main() {
	var N, A, Y, X int
	fmt.Scan(&N, &A, &X, &Y)

	res := 1000000
	if N <= A {
		res = X * N
	} else {
		res = X * A
		res += (N - A) * Y
	}

	fmt.Println(res)

}
