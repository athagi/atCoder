package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	a := N % 10
	res := "hon"
	if a == 0 || a == 1 || a == 6 || a == 8{
		res = "pon"
	} else if a == 3 {
		res = "bon"
	}

	

	fmt.Println(res)
}
