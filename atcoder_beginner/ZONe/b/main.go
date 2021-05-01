package main

import (
	"fmt"
)

func main() {
	var N, D, H int
	fmt.Scan(&N, &D, &H)

	var d, h, res, tmp  float64 = 1000, 1000, 0, 1000
	
	for i := 0; i < N; i++ {
		fmt.Scan(&d, &h)
		distance := float64(H) - h
		hight := float64(D) - d
		tmp = float64(H) - (float64(D) * distance / hight)
		
		if tmp > res {
			res = tmp
		}
	}

	fmt.Println(res)
}