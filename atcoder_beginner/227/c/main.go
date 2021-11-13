package main

import (
	"fmt"
)

func main() {
	var N int64
	fmt.Scan(&N)
	var res int64
	for a := 1; int64(a*a*a) <= N; a++ {
		for b := a; int64(a*b*b) <= N; b++ {
			c := N/int64(a*b) - int64(b) + int64(1)
			res += c
		}
	}
	fmt.Println(res)
}
