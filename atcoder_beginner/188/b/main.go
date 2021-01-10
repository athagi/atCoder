package main

import (
	"fmt"
)

func main() {
	var A, B, N int
	fmt.Scan(&N)
	As := make([]int, N)
	
	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		As[i] = A
	}

	Bs := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&B)
		Bs[i] = B
	}

	res := 0
	for i := 0; i < N; i++ {
		res += As[i] * Bs[i]
	}

	ans := "No"
	if res == 0 {
		ans = "Yes"
	}
	fmt.Println(ans)
}