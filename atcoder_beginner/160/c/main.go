package main

import (
	"fmt"
	"math"
)

func main() {
	var K, N int
	fmt.Scan(&K, &N)

	var MaxDistance float64
	var A int
	As := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A)
		As[i] = A
		if i >= 1 {
			MaxDistance = math.Max(float64(MaxDistance), float64(A-As[i-1]))
			// fmt.Println(MaxDistance, float64(As[i-1]-A))
		}
		// fmt.Println(As)
	}

	res := K - As[len(As)-1] + As[0]
	// fmt.Println(res, MaxDistance, K)
	if res < int(MaxDistance) {
		res = int(MaxDistance)
	}
	fmt.Println(K - res)

}
