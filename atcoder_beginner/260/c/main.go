package main

import "fmt"

func main() {
	var N int
	var X, Y int64
	fmt.Scan(&N, &X, &Y)
	red := make([]int64, 12)
	blue := make([]int64, 12)
	blue[1] = 1
	for n := 2; n <= N; n++ {
		blue[n] = red[n-1] + blue[n-1]*Y
		red[n] = red[n-1] + blue[n]*X

	}

	fmt.Println(red[N])
}
