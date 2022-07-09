package main

import "fmt"

func main() {
	var N, M, X, T, D int
	fmt.Scan(&N, &M, &X, &T, &D)

	height := T
	age := N
	for i := 0; i < N-M; i++ {
		if age <= X {
			height -= D
		}
		age--
	}
	fmt.Println(height)
}
