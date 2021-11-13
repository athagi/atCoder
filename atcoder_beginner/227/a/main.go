package main

import (
	"fmt"
)

func main() {
	var N, K, A int
	fmt.Scan(&N, &K, &A)
	person := make([]int, N)
	j := A
	for i := 0; i < N; i++ {
		person[i] = j
		j++
		if j > N {
			j = 1
		}
	}
	// fmt.Println(person)
	k := (K - 1) % N

	fmt.Println(person[k])

}
