package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	var a int
	res := "APPROVED"
	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		if a%2 == 0 {
			if a%3 != 0 && a%5 != 0 {
				res = "DENIED"
			}
		}
	}

	fmt.Println(res)
}
