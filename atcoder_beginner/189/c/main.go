package main

import (
		"fmt"
		"math"
)



func main() {
	var N, A int
	fmt.Scan(&N)
	As := make([]int, N)
	
	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		As[i] = A
	}

	var x float64
	var res int
	for l := 0; l < N; l++{
		x = float64(As[l])
		for r := l; r < N; r++ {
			x = math.Min(float64(x), float64(As[r]))
			res = int(math.Max(float64(res), x*float64(r-l+1)))
		}
	}
	
	fmt.Println(res)
}