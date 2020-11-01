package main

import (
	"fmt"
)


func main() {
	var N, X, Y int
	fmt.Scan(&N)
	point := make([][]int, N)
	res := "No"
	for i := 0; i < N; i++ {
		fmt.Scan(&X, &Y)
		point[i] = []int{X,Y}
	}
	var x, y, z []int
	for i:= 0; i < N; i++ {
		for j := i+1; j < N; j++ {
			for k := j+1; k < N; k++ {
				x = point[i]
				y = point[j]
				z = point[k]
				if x[0] == y[0] && y[0] == z[0] && z[0] == x[0] {
					res = "Yes"
				}
				if x[1] == y[1] && y[1] == z[1] && z[1] == x[1] {
					res = "Yes"
				}
				dx1 := x[0] - y[0]
				dx2 := x[0] - z[0]
				dy1 := x[1] - y[1]
				dy2 := x[1] - z[1]
				if dx1 * dy2 == dx2 * dy1 {
					res = "Yes"
				}
			}
		} 
	}
	
	fmt.Println(res)
}

   