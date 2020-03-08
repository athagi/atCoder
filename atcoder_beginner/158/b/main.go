package main

import (
	"fmt"
)

func main() {
	var N, A, B int64
	fmt.Scan(&N, &A, &B)

	remain := N % (A + B)
	set := (N - remain) / (A + B)

	// これだとダメ。多分あふれた
	// res := A*set + int64(math.Min(float64(remain), float64(A)))
	var res int64
	if remain > A {
		res = A*set + A
	} else {
		res = A*set + remain
	}
	fmt.Println(res)
}
