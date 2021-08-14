package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	res := 0

	if N <= 125 {
		res = 4
	} else if N <= 211 {
		res = 6
	} else if N <= 214 {
		res = 8
	}

	fmt.Println(res)

}
