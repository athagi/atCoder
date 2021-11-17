package main

import (
	"fmt"
	"math"
)

type City struct {
	X int
	Y int
}

func main() {
	var N, X, Y int
	fmt.Scan(&N)
	city := make([]City, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&X, &Y)
		city[i] = City{X: X, Y: Y}
	}

	pattern := make([]float64, combination(N, 2))
	k := 0
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			pattern[k] = dist(city[i], city[j])
			k++
		}
	}
	var res float64
	for i := 0; i < len(pattern); i++ {
		res += pattern[i]
	}
	fmt.Println(res / float64(len(pattern)) * float64(N-1))
}

func dist(a, b City) float64 {
	tmp := (a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)
	return math.Sqrt(float64(tmp))
}

func permutation(n int, k int) int {
	v := 1
	if 0 < k && k <= n {
		for i := 0; i < k; i++ {
			v *= (n - i)
		}
	} else if k > n {
		v = 0
	}
	return v
}

func factorial(n int) int {
	return permutation(n, n-1)
}

func combination(n int, k int) int {
	return permutation(n, k) / factorial(k)
}
