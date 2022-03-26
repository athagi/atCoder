package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, A int
	fmt.Scan(&N)
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		a[i] = A
	}
	sort.Ints(a)

	res := 0
	for i := 0; i < N; i++ {
		x := a[i]
		if res == x {
			res += 1
		} else if res < x {
			break
		}
	}
	fmt.Println(res)
}
