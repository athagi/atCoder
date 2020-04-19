package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	As := make([]int, N)
	var A int
	for i := 1; i < N; i++ {
		fmt.Scan(&A)
		As[A-1]++
	}

	for i := 0; i < len(As); i++ {
		fmt.Println(As[i])
	}

}
