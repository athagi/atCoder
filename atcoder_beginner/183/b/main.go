package main

import (
	"fmt"
)

func main() {
	var sx, sy, gx, gy float64
	fmt.Scan(&sx, &sy, &gx, &gy)
	
	res := ((gy * sx) + (sy * gx)) / (sy + gy)
	fmt.Println(res)

}