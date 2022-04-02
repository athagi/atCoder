package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B float64
	fmt.Scan(&A, &B)

	dist := math.Sqrt(A*A + B*B)
	a := A / dist
	b := B / dist
	fmt.Println(a, b)
}
