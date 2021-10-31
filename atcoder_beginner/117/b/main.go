package main

import (
	"fmt"
)

func main() {
	var N, l int
	fmt.Scan(&N)

	L := make([]int, N)
	max := 0
	sum := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&l)
		L[i] = l
		if max < l {
			max = l
		}
		sum += l
	}
	res := "No"
	// fmt.Println(sum, max)
	if sum-max > max {
		res = "Yes"
	}
	fmt.Println(res)
}
