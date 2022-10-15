package main

import (
	"fmt"
	"math"
)

func main() {
	var X int64
	var K int
	fmt.Scan(&X)
	fmt.Scan(&K)
	for i := 0; i < K; i++ {
		basenum := math.Pow10(i)
		x := X
		y := x / int64(basenum)
		m := y % 10
		if m <= 4 {
			X = (y - m) * int64(basenum)
		} else {
			X = (y + 10 - m) * int64(basenum)
		}
	}
	fmt.Println(X)
}
