package main

import (
	"fmt"
)

func main() {
	var N, K, M int
	fmt.Scan(&N, &K, &M)
	a := 0
	A := []int{}
	for i := 0; i < N-1; i++ {
		fmt.Scan(&a)
		A = append(A, a)
	}

	res := -1
	for an := 0; an <= K; an++ {
		sum := an
		for i := 0; i < len(A); i++ {
			sum += A[i]
		}
		average := sum / N
		if average >= M {
			res = an
			break
		}
	}

	fmt.Println(res)
}
