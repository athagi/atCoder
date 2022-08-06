package main

import "fmt"

func main() {
	var N, p int
	fmt.Scan(&N)
	P := make([]int, N)
	P[0] = 0
	for i := 1; i < N; i++ {
		fmt.Scan(&p)
		P[i] = p
	}

	// fmt.Println(P)

	parent := P[len(P)-1]
	// fmt.Println(parent)
	gen := 1
	for {
		parent = P[parent-1]
		if parent == 0 {
			break
		}
		gen++
	}
	fmt.Println(gen)
}
