package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	As := make([]int, N)
	var A int
	var sum int
	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		As[i] = A
		sum += A
	}
	threshold := float64(sum) / float64(4*M)
	count := 0
	for i := 0; i < N; i++ {
		if float64(As[i]) >= threshold {
			count++
		}
	}
	res := "No"
	if count >= M {
		res = "Yes"
	}
	fmt.Println(res)

}
