package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, d float64
	fmt.Scan(&a, &b, &d)

	angle := math.Pi * d / float64(180)
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	x := a*cos - b*sin
	y := a*sin + b*cos
	fmt.Println(x, y)
}
