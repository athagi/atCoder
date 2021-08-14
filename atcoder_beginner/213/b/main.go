package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, a int
	fmt.Scan(&N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		A[i] = a
	}

	B := make([]int, N)
	copy(B, A)
	sort.Ints(B)
	tmp := B[len(B)-2]
	res := 0
	for i := 0; i < N; i++ {
		if A[i] == tmp {
			res = i + 1
			break
		}
	}
	fmt.Println(res)
}
