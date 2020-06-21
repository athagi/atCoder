package main

import (
	"fmt"
	"math"
	// "strconv"
)

func main() {
	var A float64
	var B float64
	fmt.Scan(&A,&B)

	res := A * B
	fmt.Println(math.Floor(res))
}
