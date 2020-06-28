package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	res := 0
	for i := 1; i <= N; i++ {
		res += f(N,i)
	}
	fmt.Println(res)
}

func f(n, x int) int {
	y := n/x
	return (y*(y+1) *x)/2
}