package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	WA := make([]int, N)
	AC := make([]int, N)
	var num int
	var result string
	for i := 0; i < M; i++ {
		fmt.Scan(&num, &result)
		if result == "WA" && AC[num-1] == 0 {
			WA[num-1] = WA[num-1] + 1
		} else {
			AC[num-1] = 1
		}
	}
	var ac, wa int
	for i := 0; i < N; i++ {
		if AC[i] > 0 {
			ac++
			wa += WA[i]

		}
	}
	fmt.Println(ac, wa)
}
