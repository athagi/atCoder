package main

import (
	"fmt"
)

func main() {
	var A, B, C, D int
	fmt.Scan(&A, &B, &C, &D)

	blue := A
	red := 0
	res := 0
	diff := blue - (red * D)
	for {
		if blue <= red*D {
			break
		}
		blue += B
		red += C
		res++
		if diff <= blue-(red*D) {
			res = -1
			break
		}
	}

	fmt.Println(res)

}
