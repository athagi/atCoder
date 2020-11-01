package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	res := "White"
	if N % 2 == 1 {
		res = "Black"
	}

	fmt.Println(res)


}
