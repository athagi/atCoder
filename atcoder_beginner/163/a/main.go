package main

import (
	"fmt"
	"math"
)

func main() {
	var R int
	fmt.Scan(&R)

	res := float64(R) * 2 * math.Pi
	fmt.Println(res)
}
