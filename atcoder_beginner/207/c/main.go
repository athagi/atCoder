package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	var t, l, r int64

	input := make([][]int64, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&t, &l, &r)
		l = l * 10
		r = r * 10
		if t == 2 {
			r -= 5
		} else if t == 3 {
			l += 5
		} else if t == 4 {
			l += 5
			r -= 5
		}
		input[i] = []int64{t, l, r}
	}

	res := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			a := input[i]
			b := input[j]
			if math.Max(float64(a[1]), float64(b[1])) <= math.Min(float64(a[2]), float64(b[2])) {
				res++
				continue
			}
		}
	}

	fmt.Println(res)
}
