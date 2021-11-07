package main

import (
	"fmt"
	"math"
)

func main() {
	var X float64
	fmt.Scan(&X)
	fmt.Printf("%.0f", math.Round(X))
}
