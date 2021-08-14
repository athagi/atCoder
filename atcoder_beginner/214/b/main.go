package main

import (
	"fmt"
)

func main() {
	var S, T int
	fmt.Scan(&S, &T)

	res := 0
	for a := 0; a <= S; a++ {
		for b := 0; b <= S-a; b++ {
			for c := 0; c <= S-a-b; c++ {
				if a+b+c <= S && a*b*c <= T {
					res++
				}
			}
		}
	}

	fmt.Println(res)
}
