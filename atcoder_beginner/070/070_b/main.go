package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, d float64
	fmt.Scan(&a, &b, &c, &d)
	var res float64
	if b-c <= 0 || d-a <= 0 {
		res = 0
	} else {
		// if a <= c && b >= d {
		// 	res = d - c
		// } else if a >= c && b <= d {
		// 	res = b - a
		// } else if a >= c && b >= d {
		// 	res = d - a
		// } else {
		// 	res = b - c
		// }
		min := math.Min(b, d)
		max := math.Max(a, c)
		res = min - max
	}
	fmt.Println(res)
}
