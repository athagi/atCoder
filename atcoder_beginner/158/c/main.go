package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	var x1, x2 float64
	res := 0
	for i := 1; i <= 12500; i++ {
		x1 = math.Floor(float64(i) * 0.08)
		x2 = math.Floor(float64(i) * 0.1)
		// fmt.Println(x1, x2)

		if int(x1) == A && int(x2) == B {
			res = i
			break
		}
	}

	if res == 0 {
		res = -1
	}
	fmt.Println(res)

}
