package main

import (
	"fmt"
)

func main() {
	var N, A int
	fmt.Scan(&N)

	res := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		if A > 10 {
			res += A - 10
		}
	}

	fmt.Println(res)

}
