package main

import (
	"fmt"
)

func main() {
	var N, X, V, P int
	fmt.Scan(&N, &X)

	alc := 0
	res := -1
	for i := 0; i < N; i++ {
		fmt.Scan(&V, &P)
		alc += V * P
		if alc > X * 100 {
			res = i+1
			break;
		}
	}
	fmt.Println(res)
}