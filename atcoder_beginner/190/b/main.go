package main

import (
	"fmt"
)

func main() {
	var N, S, D, X, Y int
	fmt.Scan(&N, &S, &D)

	res := "No"
	for i := 0; i < N; i++ {
		fmt.Scan(&X, &Y)
		if S > X && D < Y {
			res = "Yes"
			break
		}
	}

	fmt.Println(res)
}