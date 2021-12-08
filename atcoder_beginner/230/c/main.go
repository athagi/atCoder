package main

import (
	"fmt"
)

type position struct {
	i int64
	j int64
}

func main() {
	var N, A, B, P, Q, R, S int64
	fmt.Scan(&N, &A, &B, &P, &Q, &R, &S)
	center := position{A, B}
	for i := P; i <= Q; i++ {
		for j := R; j <= S; j++ {
			color := "."
			if isBlack(center, position{i, j}) {
				color = "#"
			}
			fmt.Print(color)
		}
		fmt.Println("")
	}
}

func isBlack(center position, pos position) bool {
	distA := center.i - pos.i
	distB := center.j - pos.j
	if distA < 0 {
		distA = -1 * distA
	}
	if distB < 0 {
		distB = -1 * distB
	}
	res := false
	if distA == distB {
		res = true
	}
	return res

}
